package quizquestion

import (
	"context"
	qzquestionpb "quiz/gunk/v1/quiz_question"
	"quiz/usermgm/storage"
)

type CoreQuizQuestion interface {
	CreateQuizQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error)
	QuizQuestionIdForEdit(id int) (*storage.QuizQuestion, error)
	DeleteQuizQuestion(id int32) error
	QuizQuestionList(ql storage.QuizQuestionFilter) ([]storage.QzQuestion, error)
	UpdateQuiQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error)
}
type QuizQuestionSvc struct {
	qzquestionpb.UnimplementedQuizQuestionServiceServer
	core CoreQuizQuestion
}

func NewQuizQuestionSvc(qzq CoreQuizQuestion) *QuizQuestionSvc {
	return &QuizQuestionSvc{
		core: qzq,
	}
}
func (qzq QuizQuestionSvc) CreateQuizQuestion(ctx context.Context, r *qzquestionpb.CreateQuizQuestionRequest) (*qzquestionpb.CreateQuizQuestionResponse, error) {
	qquestion := storage.QuizQuestion{
		ID:         0,
		QuizID:     int(r.GetQuizID()),
		QuestionID: int(r.GetQuestionID()),
	}
	// if err := qquestion.Validate(); err != nil {
	// 	return nil, err //TODO:: will fix when implement this service in cms
	// }
	q, err := qzq.core.CreateQuizQuestion(qquestion)
	if err != nil {
		return nil, err
	}
	return &qzquestionpb.CreateQuizQuestionResponse{
		QuizQuestion: &qzquestionpb.QuizQuestion{
			ID:         int32(q.ID),
			QuizID:     int32(q.QuizID),
			QuestionID: int32(q.QuestionID),
		},
	}, nil
}
func (qzq QuizQuestionSvc) EditQuizQuestion(ctx context.Context, r *qzquestionpb.EditQuizQuestionRequest) (*qzquestionpb.EditQuizQuestionResponse, error) {
	qu, err := qzq.core.QuizQuestionIdForEdit(int(r.ID))
	if err != nil {
		return nil, err
	}
	return &qzquestionpb.EditQuizQuestionResponse{
		QuizQuestion: &qzquestionpb.QuizQuestion{
			ID:         int32(qu.ID),
			QuizID:     int32(qu.QuizID),
			QuestionID: int32(qu.QuestionID),
		},
	}, nil
}
func (qzq QuizQuestionSvc) DeleteQuizQuestion(ctx context.Context, r *qzquestionpb.DeleteQuizQuestionRequest) (*qzquestionpb.DeleteQuizQuestionResponse, error) {
	err := qzq.core.DeleteQuizQuestion(r.ID)
	if err != nil {
		return nil, err
	}
	return &qzquestionpb.DeleteQuizQuestionResponse{}, nil
}

func (qzq QuizQuestionSvc) ListQuizQuestion(ctx context.Context, r *qzquestionpb.ListQuizQuestionRequest) (*qzquestionpb.ListQuizQuestionResponse, error) {

	qzquestion := storage.QuizQuestionFilter{
		SearchTerm: r.GetSearchTerm(),
		Sortby:     int(r.GetSortBy()),
	}
	quizquestionList, err := qzq.core.QuizQuestionList(qzquestion)
	if err != nil {
		return nil, err
	}

	var totalQuizQuestion []*qzquestionpb.ListQzQuestion
	for _, qtl := range quizquestionList {
		user := &qzquestionpb.ListQzQuestion{
			ID:           int32(qtl.ID),
			Title:        qtl.Title,
			CategoryName: qtl.CategoryName.String,
			QuizTitle:    qtl.QuizTitle,
			QuizTime:     int32(qtl.QuizTime),
		}
		totalQuizQuestion = append(totalQuizQuestion, user)
	}
	return &qzquestionpb.ListQuizQuestionResponse{
		QuizQuestionFilterList: &qzquestionpb.QuizQuestionFilterList{
			TotalQuizQuestion: totalQuizQuestion,
			SearchTerm:        "",
			SortBy:            0,
		},
	}, nil
	// question := make([]*qzquestionpb.ListQzQuestion, len(quizquestionList))
	// for i, ql := range quizquestionList {
	// 	question[i] = &qzquestionpb.ListQzQuestion{
	// 		ID:           int32(ql.ID),
	// 		Title:        ql.Title,
	// 		CategoryName: ql.CategoryName,
	// 		QuizTitle:    ql.QuizTitle,
	// 		QuizTime:     int32(ql.QuizTime),
	// 	}
	// }
	// return &qzquestionpb.ListQuizQuestionResponse{
	// 	ListQzQuestion: question,
	// }, nil
}
func (qzq QuizQuestionSvc) UpdateQuizQuestion(ctx context.Context, r *qzquestionpb.UpdateQuizQuestionRequest) (*qzquestionpb.UpdateQuizQuestionResponse, error) {
	question := storage.QuizQuestion{
		ID:         int(r.ID),
		QuizID:     int(r.QuizID),
		QuestionID: int(r.QuestionID),
	}
	_, err := qzq.core.UpdateQuiQuestion(question)
	if err != nil {
		return nil, err
	}
	return &qzquestionpb.UpdateQuizQuestionResponse{}, nil
}
