package quiz

import (
	"fmt"
	"quiz/usermgm/storage"
)

type QuizStore interface {
	CreateQuiz(u storage.Quiz) (*storage.Quiz, error)
	DeleteQuiz(id int32) error
	QuizIdForEdit(id int) (*storage.Quiz, error)
	UpdateQuiz(u storage.Quiz) (*storage.Quiz, error)
	QuizList(qq storage.QuizFilter) ([]storage.QuizList, error)
}
type CoreQuiz struct {
	store QuizStore
}

func NewCoreQuiz(qz QuizStore) *CoreQuiz {
	return &CoreQuiz{
		store: qz,
	}
}
func (qz CoreQuiz) CreateQuiz(u storage.Quiz) (*storage.Quiz, error) {

	qu, err := qz.store.CreateQuiz(u)
	if err != nil {
		return nil, err
	}
	if qu == nil {
		return nil, fmt.Errorf("unable to create quiz")
	}
	return qu, nil
}

func (qz CoreQuiz) DeleteQuiz(id int32) error {
	err := qz.store.DeleteQuiz(id)
	if err != nil {
		return err
	}
	return nil
}

func (qz CoreQuiz) QuizIdForEdit(id int) (*storage.Quiz, error) {
	quiz, err := qz.store.QuizIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if quiz == nil {
		return nil, err
	}
	return quiz, nil
}
func (qz CoreQuiz) UpdateQuiz(u storage.Quiz) (*storage.Quiz, error) {
	quiz, err := qz.store.UpdateQuiz(u)
	if err != nil {
		return nil, err
	}
	if quiz == nil {
		return nil, err
	}
	return quiz, nil
}
func (qz CoreQuiz) QuizList(qq storage.QuizFilter) ([]storage.QuizList, error) {
	list, err := qz.store.QuizList(qq)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, fmt.Errorf("unable to show quiz list")
	}

	return list, nil
}
