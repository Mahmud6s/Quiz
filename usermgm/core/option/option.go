package option

import (
	"fmt"
	"quiz/usermgm/storage"
)

type OptionStore interface {
	CreateOption(u storage.Option) (*storage.Option, error)
	DeleteOption(id int32) error
	GetOptionIdForEdit(id int) (*storage.Option, error)
	UpdateOption(u storage.Option) (*storage.Option, error)
	OptionList(op storage.OptionsFilter) ([]storage.Option, error)
}
type CoreOption struct {
	store OptionStore
}

func NewCoreOption(opt OptionStore) *CoreOption {
	return &CoreOption{
		store: opt,
	}
}
func (opt CoreOption) CreateOption(u storage.Option) (*storage.Option, error) {
	oc, err := opt.store.CreateOption(u)
	if err != nil {
		return nil, err
	}
	if oc == nil {
		return nil, fmt.Errorf("unable to create option")
	}
	return oc, nil
}
func (opt CoreOption) DeleteOption(id int32) error {
	err := opt.store.DeleteOption(id)
	if err != nil {
		return err
	}
	return nil
}
func (opt CoreOption) GetOptionIdForEdit(id int) (*storage.Option, error) {
	option, err := opt.store.GetOptionIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if option == nil {
		return nil, err
	}
	return option, nil
}
func (opt CoreOption) UpdateOption(u storage.Option) (*storage.Option, error) {
	option, err := opt.store.UpdateOption(u)
	if err != nil {
		return nil, err
	}
	if option == nil {
		return nil, err
	}
	return option, nil
}
func (opt CoreOption) OptionList(op storage.OptionsFilter) ([]storage.Option, error) {
	optionlist, err := opt.store.OptionList(op)
	if err != nil {
		return nil, err
	}
	if optionlist == nil {
		return nil, fmt.Errorf("unable to show option list")
	}

	return optionlist, nil
}
