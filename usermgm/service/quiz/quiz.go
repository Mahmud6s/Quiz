package quiz

import (
	"context"
	"quiz/usermgm/storage"

	quizpb "quiz/gunk/v1/quiz"
)

type CoreQuiz interface {
	CreateQuiz(u storage.Quiz) (*storage.Quiz, error)
	DeleteQuiz(id int32) error
	QuizIdForEdit(id int) (*storage.Quiz, error)
	UpdateQuiz(u storage.Quiz) (*storage.Quiz, error)
	QuizList(qq storage.QuizFilter) ([]storage.QuizList, error)
}
type QuizSvc struct {
	quizpb.UnimplementedQuizServiceServer
	core CoreQuiz
}

func NewQuizSvc(qz CoreQuiz) *QuizSvc {
	return &QuizSvc{
		core: qz,
	}
}

func (qz QuizSvc) CreateQuiz(ctx context.Context, r *quizpb.CreateQuizRequest) (*quizpb.CreateQuizResponse, error) {
	quiz := storage.Quiz{
		ID:         0,
		CategoryID: int(r.GetCategoryID()),
		QuizTitle:  r.GetQuizTitle(),
		QuizTime:   int(r.QuizTime),
	}
	if err := quiz.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}
	c, err := qz.core.CreateQuiz(quiz)
	if err != nil {
		return nil, err
	}
	return &quizpb.CreateQuizResponse{
		Quiz: &quizpb.Quiz{
			ID:         int32(c.ID),
			CategoryID: int32(c.CategoryID),
			QuizTitle:  c.QuizTitle,
			QuizTime:   int32(c.QuizTime),
		},
	}, nil
}
func (qz QuizSvc) DeleteQuiz(ctx context.Context, r *quizpb.DeleteQuizRequest) (*quizpb.DeleteQuizResponse, error) {
	err := qz.core.DeleteQuiz(r.ID)
	if err != nil {
		return nil, err
	}
	return &quizpb.DeleteQuizResponse{}, nil
}

func (qz QuizSvc) EditQuiz(ctx context.Context, r *quizpb.EditQuizRequest) (*quizpb.EditQuizResponse, error) {
	rq, err := qz.core.QuizIdForEdit(int(r.ID))
	if err != nil {
		return nil, err
	}
	return &quizpb.EditQuizResponse{
		Quiz: &quizpb.Quiz{
			ID:         int32(rq.ID),
			CategoryID: int32(rq.CategoryID),
			QuizTitle:  rq.QuizTitle,
			QuizTime:   int32(rq.QuizTime),
		},
	}, nil
}
func (qz QuizSvc) UpdateQuiz(ctx context.Context, r *quizpb.UpdateQuizRequest) (*quizpb.UpdateQuizResponse, error) {
	quiz := storage.Quiz{
		ID:         int(r.ID),
		CategoryID: int(r.CategoryID),
		QuizTitle:  r.QuizTitle,
		QuizTime:   int(r.QuizTime),
	}
	_, err := qz.core.UpdateQuiz(quiz)
	if err != nil {
		return nil, err
	}
	return &quizpb.UpdateQuizResponse{}, nil
}
func (qz QuizSvc) ListQuiz(ctx context.Context, r *quizpb.ListQuizRequest) (*quizpb.ListQuizResponse, error) {
	qquiz := storage.QuizFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	quizListt, err := qz.core.QuizList(qquiz)
	if err != nil {
		return nil, err
	}
	var totalQuiz []*quizpb.ListQuiz
	for _, qtl := range quizListt {
		user := &quizpb.ListQuiz{
			ID:           int32(qtl.ID),
			CategoryName: qtl.CategoryName,
			QuizTitle:    qtl.QuizTitle,
		}
		totalQuiz = append(totalQuiz, user)
	}
	return &quizpb.ListQuizResponse{
		QuizFilterList: &quizpb.QuizFilterList{
			TotalQuiz:  totalQuiz,
			SearchTerm: "",
		},
	}, nil
}
