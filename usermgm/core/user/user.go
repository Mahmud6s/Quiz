package user

import (
	"fmt"
	"quiz/usermgm/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	RegisterUser(storage.User) (*storage.User, error)
	RegisterAdmin(storage.User) (*storage.User, error)
	GetUserByUsername(string) (*storage.User, error)
	UserList() ([]storage.User, error)
	DeleteUser(id int32) error
	GetUserIdForEdit(id int) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
}

type CoreUser struct {
	store UserStore
}

func NewCoreUser(us UserStore) *CoreUser {
	return &CoreUser{
		store: us,
	}
}

func (cu CoreUser) RegisterUser(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashPass)
	ru, err := cu.store.RegisterUser(u)
	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return ru, nil
}

func (cu CoreUser) RegisterAdmin(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u.Password = string(hashPass)
	ru, err := cu.store.RegisterAdmin(u)
	if err != nil {
		return nil, err
	}

	if ru == nil {
		return nil, fmt.Errorf("unable to register")
	}

	return ru, nil
}

func (cu CoreUser) UserList() ([]storage.User, error) {
	userlist, err := cu.store.UserList()
	if err != nil {
		return nil, err
	}
	if userlist == nil {
		return nil, fmt.Errorf("unable to show user list")
	}

	return userlist, nil
}
func (cu CoreUser) DeleteUser(id int32) error {
	err := cu.store.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
func (cu CoreUser) GetUserIdForEdit(id int) (*storage.User, error) {

	user, err := cu.store.GetUserIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	return user, nil
}
func (cu CoreUser) UpdateUser(u storage.User) (*storage.User, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashPass)
	user, err := cu.store.UpdateUser(u)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, err
	}
	return user, nil
}
func (cu CoreUser) Login(l storage.Login) (*storage.User, error) {
	u, err := cu.store.GetUserByUsername(l.Username)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		return nil, err
	}

	return u, nil
}
