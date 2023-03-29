package handler

import (
	"log"
	"net/http"
	userpb "quiz/gunk/v1/user"
)

type List struct {
	Users []User
}
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
	IsActive  bool
	IsAdmin   bool
}

func (h Handler) UserListR(w http.ResponseWriter, r *http.Request) {
	list, err := h.usermgmSvc.UserList(r.Context(), &userpb.UserListRequest{})
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
	data := []User{}
	if list != nil {
		for _, u := range list.User {
			data = append(data, User{
				ID:        int(u.ID),
				FirstName: u.FirstName,
				LastName:  u.LastName,
				Email:     u.Email,
				Username:  u.Username,
				IsActive:  u.IsActive,
				IsAdmin:   u.IsAdmin,
			})
		}
	}
	h.parseUserListTemplate(w, List{
		Users: data,
	})
}

func (h Handler) parseUserListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("userList.html")
	if t == nil {
		log.Println("unable to lookup create list template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
