package handler

import (
	"log"
	"net/http"
	"strings"

	userpb "quiz/gunk/v1/user"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/justinas/nosurf"
)

type RegisterUser struct {
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

func (u RegisterUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.Error("The username field is required."),
		),
		validation.Field(&u.Email,
			validation.Required.Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.Error("The password field is required."),
		),
	)
}

type RegisterUserForm struct {
	RegisterUser RegisterUser
	FormError    map[string]error
	CSRFToken    string
}

func (h Handler) Register(w http.ResponseWriter, r *http.Request) {
	h.pareseRegisterTemplate(w, RegisterUserForm{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) pareseRegisterTemplate(w http.ResponseWriter, form RegisterUserForm) {
	t := h.Templates.Lookup("register.html")
	if t == nil {
		log.Println("unable to lookup register template")
		h.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, form); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) RegisterPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	form := RegisterUserForm{}
	ru := RegisterUser{}
	if err := h.decoder.Decode(&ru, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	form.RegisterUser = ru
	if err := ru.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			for key, val := range vErr {
				form.FormError[strings.Title(key)] = val
			}
		}
		h.pareseRegisterTemplate(w, form)
		return
	}

	_, err := h.usermgmSvc.RegisterUser(r.Context(), &userpb.RegisterUserRequest{
		FirstName: ru.FirstName,
		LastName:  ru.LastName,
		Username:  ru.Username,
		Email:     ru.Email,
		Password:  ru.Password,
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
