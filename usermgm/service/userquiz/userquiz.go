package userquiz

import (
	"context"
	userquizpb "quiz/gunk/v1/user_quiz"
	"quiz/usermgm/storage"
)

type CoreUserQuiz interface {
	CreateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error)
	UserQuizIdForEdit(id int) (*storage.UserQuiz, error)
	DeleteUserQuiz(id int32) error
	UpdateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error)
	UserQuizList(uqq storage.UserQuizFilter) ([]storage.UserQuizList, error)
}
type UserQuizSvc struct {
	userquizpb.UnimplementedUserQuizServiceServer
	core CoreUserQuiz
}

func NewUserQuizSvc(uq CoreUserQuiz) *UserQuizSvc {
	return &UserQuizSvc{
		core: uq,
	}
}
func (uq UserQuizSvc) CreateUserQuiz(ctx context.Context, r *userquizpb.CreateUserQuizRequest) (*userquizpb.CreateUserQuizResponse, error) {
	quiz := storage.UserQuiz{
		ID:             0,
		UsersID:        int(r.UserID),
		QuizQuestionID: int(r.QuizQuestionID),
	}
	// if err := quiz.Validate(); err != nil {
	// 	return nil, err //TODO:: will fix when implement this service in cms
	// }
	q, err := uq.core.CreateUserQuiz(quiz)
	if err != nil {
		return nil, err
	}
	return &userquizpb.CreateUserQuizResponse{
		UserQuiz: &userquizpb.UserQuiz{
			ID:             int32(q.ID),
			UserID:         int32(q.UsersID),
			QuizQuestionID: int32(q.QuizQuestionID),
		},
	}, nil
}
func (uq UserQuizSvc) EditUserQuiz(ctx context.Context, r *userquizpb.EditUserQuizRequest) (*userquizpb.EditUserQuizResponse, error) {
	uqz, err := uq.core.UserQuizIdForEdit(int(r.ID))
	if err != nil {
		return nil, err
	}
	return &userquizpb.EditUserQuizResponse{
		UserQuiz: &userquizpb.UserQuiz{
			ID:             int32(uqz.ID),
			UserID:         int32(uqz.UsersID),
			QuizQuestionID: int32(uqz.QuizQuestionID),
		},
	}, nil
}
func (uq UserQuizSvc) DeleteUserQuiz(ctx context.Context, r *userquizpb.DeleteUserQuizRequest) (*userquizpb.DeleteUserQuizResponse, error) {
	err := uq.core.DeleteUserQuiz(r.ID)
	if err != nil {
		return nil, err
	}
	return &userquizpb.DeleteUserQuizResponse{}, nil
}
func (uq UserQuizSvc) UpdateUserQuiz(ctx context.Context, r *userquizpb.UpdateUserQuizRequest) (*userquizpb.UpdateUserQuizResponse, error) {
	quiz := storage.UserQuiz{
		ID:             int(r.ID),
		UsersID:        int(r.UserID),
		QuizQuestionID: int(r.QuizQuestionID),
	}
	_, err := uq.core.UpdateUserQuiz(quiz)
	if err != nil {
		return nil, err
	}
	return &userquizpb.UpdateUserQuizResponse{}, nil
}

func (uq UserQuizSvc) ListUserQuiz(ctx context.Context, r *userquizpb.ListUserQuizRequest) (*userquizpb.ListUserQuizResponse, error) {
	userQZ := storage.UserQuizFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	userQZList, err := uq.core.UserQuizList(userQZ)
	if err != nil {
		return nil, err
	}
	var totalUserQZ []*userquizpb.ListUserQuiz
	for _, qzl := range userQZList {
		user := &userquizpb.ListUserQuiz{
			FirstName: qzl.FirstName,
			LastName:  qzl.LastName,
			Username:  qzl.Username,
			Title:     qzl.Title,
			QuizTitle: qzl.QuizTitle,
		}
		totalUserQZ = append(totalUserQZ, user)
	}
	return &userquizpb.ListUserQuizResponse{
		ListUserQuizFilter: &userquizpb.ListUserQuizFilter{
			TotalUserQuiz: totalUserQZ,
			SearchTerm:    "",
		},
	}, nil
}
