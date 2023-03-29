package quizquestion

import (
	"fmt"
	"quiz/usermgm/storage"
)

type QuizQuestionStore interface {
	CreateQuizQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error)
	QuizQuestionIdForEdit(id int) (*storage.QuizQuestion, error)
	DeleteQuizQuestion(id int32) error
	QuizQuestionList(ql storage.QuizQuestionFilter) ([]storage.QzQuestion, error)
	UpdateQuiQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error)
}
type CoreQuizQuestion struct {
	store QuizQuestionStore
}

func NewCoreQuestion(qzq QuizQuestionStore) *CoreQuizQuestion {
	return &CoreQuizQuestion{
		store: qzq,
	}
}

func (qzq CoreQuizQuestion) CreateQuizQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error) {
	qz, err := qzq.store.CreateQuizQuestion(u)
	if err != nil {
		return nil, err
	}
	if qz == nil {
		return nil, fmt.Errorf("unable to create quiz_question")
	}
	return qz, nil
}
func (qzq CoreQuizQuestion) QuizQuestionIdForEdit(id int) (*storage.QuizQuestion, error) {
	qquestion, err := qzq.store.QuizQuestionIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if qquestion == nil {
		return nil, err
	}
	return qquestion, nil
}
func (qzq CoreQuizQuestion) DeleteQuizQuestion(id int32) error {
	err := qzq.store.DeleteQuizQuestion(id)
	if err != nil {
		return err
	}
	return nil
}

func (qzq CoreQuizQuestion) QuizQuestionList(ql storage.QuizQuestionFilter) ([]storage.QzQuestion, error) {
	list, err := qzq.store.QuizQuestionList(ql)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return nil, fmt.Errorf("unable to show quiz_question list")
	}

	return list, nil
}
func (qzq CoreQuizQuestion) UpdateQuiQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error) {
	question, err := qzq.store.UpdateQuiQuestion(u)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, err
	}
	return question, nil
}
