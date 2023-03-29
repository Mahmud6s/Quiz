package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	userpb "quiz/gunk/v1/user"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

type LoginUser struct {
	Username  string
	Password  string
	FormError map[string]error
	CSRFToken string
}

func (lu LoginUser) Validate() error {
	return validation.ValidateStruct(&lu,
		validation.Field(&lu.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&lu.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.pareseLoginTemplate(w, LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	var lf LoginUser
	if err := h.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := lf.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			formErr := make(map[string]error)
			for key, val := range vErr {
				formErr[strings.Title(key)] = val
			}
			lf.FormError = formErr
			lf.Password = ""
			lf.CSRFToken = nosurf.Token(r)
			h.pareseLoginTemplate(w, lf)
			return
		}
	}

	u, err := h.usermgmSvc.Login(r.Context(), &userpb.LoginRequest{
		Username: lf.Username,
		Password: lf.Password,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(int(u.GetUser().ID)))
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func (h Handler) pareseLoginTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		log.Println("unaable to lookup login template")
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
