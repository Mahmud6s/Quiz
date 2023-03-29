package option

import (
	"context"
	optionpb "quiz/gunk/v1/option"
	"quiz/usermgm/storage"
)

type CoreOption interface {
	CreateOption(u storage.Option) (*storage.Option, error)
	DeleteOption(id int32) error
	GetOptionIdForEdit(id int) (*storage.Option, error)
	UpdateOption(u storage.Option) (*storage.Option, error)
	OptionList(op storage.OptionsFilter) ([]storage.Option, error)
}
type OptionSvc struct {
	optionpb.UnimplementedOptionServiceServer
	core CoreOption
}

func NewOptionSvc(op CoreOption) *OptionSvc {
	return &OptionSvc{
		core: op,
	}
}
func (op OptionSvc) CreateOption(ctx context.Context, r *optionpb.CreateOptionRequest) (*optionpb.CreateOptionResponse, error) {
	option := storage.Option{
		ID:         0,
		QuestionID: int(r.GetQuestionID()),
		OptionName: r.GetOptionName(),
	}

	if err := option.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}

	u, err := op.core.CreateOption(option)
	if err != nil {
		return nil, err
	}

	return &optionpb.CreateOptionResponse{
		Option: &optionpb.Option{
			ID:         int32(u.ID),
			QuestionID: int32(u.QuestionID),
			OptionName: u.OptionName,
			IsCorrect:  u.IsCorrect,
		},
	}, nil
}
func (op OptionSvc) DeleteOption(ctx context.Context, r *optionpb.DeleteOptionRequest) (*optionpb.DeleteOptionResponse, error) {
	err := op.core.DeleteOption(r.ID)
	if err != nil {
		return nil, err
	}
	return &optionpb.DeleteOptionResponse{}, nil
}
func (opt OptionSvc) EditOption(ctx context.Context, r *optionpb.EditOptionRequest) (*optionpb.EditOptionResponse, error) {
	ou, err := opt.core.GetOptionIdForEdit(int(r.Id))
	if err != nil {
		return nil, err
	}

	return &optionpb.EditOptionResponse{
		Option: &optionpb.Option{
			ID:         int32(ou.ID),
			QuestionID: int32(ou.QuestionID),
			OptionName: ou.OptionName,
			IsCorrect:  ou.IsCorrect,
		},
	}, err
}
func (opt OptionSvc) UpdateOption(ctx context.Context, r *optionpb.UpdateOptionRequest) (*optionpb.UpdateOptionResponse, error) {
	option := storage.Option{
		ID:         int(r.GetID()),
		QuestionID: int(r.GetQuestionID()),
		OptionName: r.GetOptionName(),
		IsCorrect:  r.GetIsCorrect(),
	}

	_, err := opt.core.UpdateOption(option)
	if err != nil {
		return nil, err
	}
	return &optionpb.UpdateOptionResponse{}, nil
}
func (opt OptionSvc) ListOption(ctx context.Context, r *optionpb.ListOptionRequest) (*optionpb.ListOptionResponse, error) {
	option := storage.OptionsFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	u, err := opt.core.OptionList(option)
	if err != nil {
		return nil, err
	}

	var totalOption []*optionpb.Option
	for _, op := range u {
		user := &optionpb.Option{
			ID:         int32(op.ID),
			QuestionID: int32(op.QuestionID),
			OptionName: op.OptionName,
		}
		totalOption = append(totalOption, user)
	}
	return &optionpb.ListOptionResponse{
		OptionFilterList: &optionpb.OptionFilterList{
			TotalOption: totalOption,
			SearchTerm:  "",
		},
	}, nil
}
