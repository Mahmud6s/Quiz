package answer

import (
	"fmt"
	"quiz/usermgm/storage"
)

type AnswerStore interface {
	CreateAnswer(u storage.Answer) (*storage.Answer, error)
	AnswerIdForEdit(id int) (*storage.Answer, error)
	UpdateAnswer(u storage.Answer) (*storage.Answer, error)
	AnswerDelete(id int32) error
	AnswerList(as storage.AnswerFilter) ([]storage.AnserList, error)
}
type CoreAnswer struct {
	store AnswerStore
}

func NewCoreAnswer(aa AnswerStore) *CoreAnswer {
	return &CoreAnswer{
		store: aa,
	}
}
func (aa CoreAnswer) CreateAnswer(u storage.Answer) (*storage.Answer, error) {

	as, err := aa.store.CreateAnswer(u)
	if err != nil {
		return nil, err
	}
	if as == nil {
		return nil, fmt.Errorf("unable to create answer")
	}
	return as, nil
}
func (aa CoreAnswer) AnswerIdForEdit(id int) (*storage.Answer, error) {
	answer, err := aa.store.AnswerIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if answer == nil {
		return nil, err
	}
	return answer, nil
}
func (aa CoreAnswer) UpdateAnswer(u storage.Answer) (*storage.Answer, error) {
	answer, err := aa.store.UpdateAnswer(u)
	if err != nil {
		return nil, err
	}
	if answer == nil {
		return nil, err
	}
	return answer, nil
}
func (aa CoreAnswer) AnswerDelete(id int32) error {
	err := aa.store.AnswerDelete(id)
	if err != nil {
		return err
	}
	return nil
}
func (aa CoreAnswer) AnswerList(as storage.AnswerFilter) ([]storage.AnserList, error) {
	list, err := aa.store.AnswerList(as)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, fmt.Errorf("unable to show answer list")
	}

	return list, nil
}
