package answer

import (
	"context"
	answerpb "quiz/gunk/v1/answer"
	"quiz/usermgm/storage"
)

type CoreAnswer interface {
	CreateAnswer(u storage.Answer) (*storage.Answer, error)
	AnswerIdForEdit(id int) (*storage.Answer, error)
	UpdateAnswer(u storage.Answer) (*storage.Answer, error)
	AnswerDelete(id int32) error
	AnswerList(as storage.AnswerFilter) ([]storage.AnserList, error)
}
type AnswerSvc struct {
	answerpb.UnimplementedAnswerServiceServer
	core CoreAnswer
}

func NewAnswerSvc(aa CoreAnswer) *AnswerSvc {
	return &AnswerSvc{
		core: aa,
	}
}

func (aa AnswerSvc) CreateAnswer(ctx context.Context, r *answerpb.CreateAnswerRequest) (*answerpb.CreateAnswerResponse, error) {
	answer := storage.Answer{
		ID:         0,
		UserquizID: int(r.UserquizID),
		QuestionID: int(r.QuestionID),
		OptionID:   int(r.OptionID),
	}
	// if err := question.Validate(); err != nil {
	// 	return nil, err //TODO:: will fix when implement this service in cms
	// }
	a, err := aa.core.CreateAnswer(answer)
	if err != nil {
		return nil, err
	}
	return &answerpb.CreateAnswerResponse{
		Answer: &answerpb.Answer{
			ID:         int32(a.ID),
			UserquizID: int32(a.UserquizID),
			QuestionID: int32(a.QuestionID),
			OptionID:   int32(a.OptionID),
			IsCorrect:  a.IsCorrect,
		},
	}, nil
}
func (aa AnswerSvc) EditAnswer(ctx context.Context, r *answerpb.EditAnswerRequest) (*answerpb.EditAnswerResponse, error) {
	as, err := aa.core.AnswerIdForEdit(int(r.ID))
	if err != nil {
		return nil, err
	}
	return &answerpb.EditAnswerResponse{
		Answer: &answerpb.Answer{
			ID:         int32(as.ID),
			UserquizID: int32(as.UserquizID),
			QuestionID: int32(as.QuestionID),
			OptionID:   int32(as.OptionID),
			IsCorrect:  as.IsCorrect,
		},
	}, nil
}
func (aa AnswerSvc) UpdateAnswer(ctx context.Context, r *answerpb.UpdateAnswerRequest) (*answerpb.UpdateAnswerResponse, error) {
	answer := storage.Answer{
		ID:         int(r.ID),
		UserquizID: int(r.UserquizID),
		QuestionID: int(r.QuestionID),
		OptionID:   int(r.OptionID),
		IsCorrect:  r.IsCorrect,
	}
	_, err := aa.core.UpdateAnswer(answer)
	if err != nil {
		return nil, err
	}
	return &answerpb.UpdateAnswerResponse{}, nil
}
func (aa AnswerSvc) DeleteAnswer(ctx context.Context, r *answerpb.DeleteAnswerRequest) (*answerpb.DeleteAnswerResponse, error) {
	err := aa.core.AnswerDelete(r.ID)
	if err != nil {
		return nil, err
	}
	return &answerpb.DeleteAnswerResponse{}, nil
}
func (aa AnswerSvc) AnswerList(ctx context.Context, r *answerpb.AnswerListRequest) (*answerpb.AnswerListResponse, error) {
	answer := storage.AnswerFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	quizquestionList, err := aa.core.AnswerList(answer)
	if err != nil {
		return nil, err
	}

	var TotalAnswer []*answerpb.AnswerList
	for _, aas := range quizquestionList {
		user := &answerpb.AnswerList{
			ID:         int32(aas.ID),
			FirstName:  aas.FirstName,
			LastName:   aas.LastName,
			Username:   aas.Username,
			Title:      aas.Title,
			OptionName: aas.OptionName,
			IsCorrect:  aas.IsCorrect,
		}
		TotalAnswer = append(TotalAnswer, user)
	}
	return &answerpb.AnswerListResponse{
		AnswerListFilter: &answerpb.AnswerListFilter{
			TotalAnswer: TotalAnswer,
			SearchTerm:  "",
		},
	}, nil
}
