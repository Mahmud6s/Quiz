package question

import (
	"context"
	questionpb "quiz/gunk/v1/question"
	"quiz/usermgm/storage"
)

type CoreQuestion interface {
	CreateQuestion(u storage.Question) (*storage.Question, error)
	DeleteQuestion(id int32) error
	QuestionIdForEdit(id int) (*storage.Question, error)
	UpdateQuestion(u storage.Question) (*storage.Question, error)
	QuestionList(qs storage.QuestionFilter) ([]storage.QuestionList, error)
}
type QuestionSvc struct {
	questionpb.UnimplementedQuestionServiceServer
	core CoreQuestion
}

func NewQuestionSvc(qq CoreQuestion) *QuestionSvc {
	return &QuestionSvc{
		core: qq,
	}
}
func (qq QuestionSvc) CreateQuestion(ctx context.Context, r *questionpb.CreateQuestionRequest) (*questionpb.CreateQuestionResponse, error) {
	question := storage.Question{
		ID:         0,
		CategoryID: int(r.GetCategoryID()),
		Title:      r.GetTitle(),
	}
	if err := question.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}
	c, err := qq.core.CreateQuestion(question)
	if err != nil {
		return nil, err
	}
	return &questionpb.CreateQuestionResponse{
		Question: &questionpb.Question{
			ID:         int32(c.ID),
			CategoryID: int32(c.CategoryID),
			Title:      c.Title,
		},
	}, nil
}

func (qq QuestionSvc) DeleteQuestion(ctx context.Context, r *questionpb.DeleteQuestionRequest) (*questionpb.DeleteQuestionResponse, error) {
	err := qq.core.DeleteQuestion(r.ID)
	if err != nil {
		return nil, err
	}
	return &questionpb.DeleteQuestionResponse{}, nil
}
func (qq QuestionSvc) EditQuestion(ctx context.Context, r *questionpb.EditQuestionRequest) (*questionpb.EditQuestionResponse, error) {
	ru, err := qq.core.QuestionIdForEdit(int(r.Id))
	if err != nil {
		return nil, err
	}

	return &questionpb.EditQuestionResponse{
		Question: &questionpb.Question{
			ID:         int32(ru.ID),
			CategoryID: int32(ru.CategoryID),
			Title:      ru.Title,
		},
	}, err
}

func (qq QuestionSvc) UpdateQuestion(ctx context.Context, r *questionpb.UpdateQuestionRequest) (*questionpb.UpdateQuestionResponse, error) {
	question := storage.Question{
		ID:         int(r.ID),
		CategoryID: int(r.CategoryID),
		Title:      r.Title,
	}
	_, err := qq.core.UpdateQuestion(question)
	if err != nil {
		return nil, err
	}
	return &questionpb.UpdateQuestionResponse{}, nil
}
func (qq QuestionSvc) ListQuestion(ctx context.Context, r *questionpb.ListQuestionRequest) (*questionpb.ListQuestionResponse, error) {
	qquestion := storage.QuestionFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	questionList, err := qq.core.QuestionList(qquestion)
	if err != nil {
		return nil, err
	}
	var totalQuestion []*questionpb.ListQuestion

	for _, qtl := range questionList {
		user := &questionpb.ListQuestion{
			ID:           int32(qtl.ID),
			CategoryName: qtl.CategoryName,
			Title:        qtl.Title,
		}
		totalQuestion = append(totalQuestion, user)
	}
	return &questionpb.ListQuestionResponse{
		QuestionFilterList: &questionpb.QuestionFilterList{
			TotalQuestion: totalQuestion,
			SearchTerm:    "",
		},
	}, nil
}
