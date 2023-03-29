package userquiz

import (
	"fmt"
	"quiz/usermgm/storage"
)

type UserQuizStore interface {
	CreateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error)
	UserQuizIdForEdit(id int) (*storage.UserQuiz, error)
	DeleteUserQuiz(id int32) error
	UpdateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error)
	UserQuizList(uqq storage.UserQuizFilter) ([]storage.UserQuizList, error)
}
type CoreUserQuiz struct {
	store UserQuizStore
}

func NewCoreUserQuiz(uq UserQuizStore) *CoreUserQuiz {
	return &CoreUserQuiz{
		store: uq,
	}
}
func (uq CoreUserQuiz) CreateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error) {
	qu, err := uq.store.CreateUserQuiz(u)
	if err != nil {
		return nil, err
	}
	if qu == nil {
		return nil, fmt.Errorf("unable to create user_quiz")
	}
	return qu, nil
}
func (uq CoreUserQuiz) UserQuizIdForEdit(id int) (*storage.UserQuiz, error) {
	userquiz, err := uq.store.UserQuizIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if userquiz == nil {
		return nil, err
	}
	return userquiz, nil
}
func (uq CoreUserQuiz) DeleteUserQuiz(id int32) error {
	err := uq.store.DeleteUserQuiz(id)
	if err != nil {
		return err
	}
	return nil
}
func (uq CoreUserQuiz) UpdateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error) {
	quiz, err := uq.store.UpdateUserQuiz(u)
	if err != nil {
		return nil, err
	}
	if quiz == nil {
		return nil, err
	}
	return quiz, nil
}
func (uq CoreUserQuiz) UserQuizList(uqq storage.UserQuizFilter) ([]storage.UserQuizList, error) {
	list, err := uq.store.UserQuizList(uqq)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, fmt.Errorf("unable to show UserQuiz list")
	}

	return list, nil
}
