package question

import (
	"fmt"
	"quiz/usermgm/storage"
)

type QuestionStore interface {
	CreateQuestion(u storage.Question) (*storage.Question, error)
	DeleteQuestion(id int32) error
	QuestionIdForEdit(id int) (*storage.Question, error)
	UpdateQuestion(u storage.Question) (*storage.Question, error)
	QuestionList(qs storage.QuestionFilter) ([]storage.QuestionList, error)
}
type CoreQuestion struct {
	store QuestionStore
}

func NewCoreQuestion(qq QuestionStore) *CoreQuestion {
	return &CoreQuestion{
		store: qq,
	}
}
func (qq CoreQuestion) CreateQuestion(u storage.Question) (*storage.Question, error) {

	qu, err := qq.store.CreateQuestion(u)
	if err != nil {
		return nil, err
	}
	if qu == nil {
		return nil, fmt.Errorf("unable to create question")
	}
	return qu, nil
}

func (qq CoreQuestion) DeleteQuestion(id int32) error {
	err := qq.store.DeleteQuestion(id)
	if err != nil {
		return err
	}
	return nil
}
func (qq CoreQuestion) QuestionIdForEdit(id int) (*storage.Question, error) {
	question, err := qq.store.QuestionIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, err
	}
	return question, nil
}
func (qq CoreQuestion) UpdateQuestion(u storage.Question) (*storage.Question, error) {
	question, err := qq.store.UpdateQuestion(u)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, err
	}
	return question, nil
}
func (qq CoreQuestion) QuestionList(qs storage.QuestionFilter) ([]storage.QuestionList, error) {
	list, err := qq.store.QuestionList(qs)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, fmt.Errorf("unable to show question list")
	}

	return list, nil
}
