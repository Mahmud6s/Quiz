package category

import (
	"fmt"
	"quiz/usermgm/storage"
)

type CategoryStore interface {
	CreateCategory(u storage.Category) (*storage.Category, error)
	DeleteCategory(id int32) error
	CategoryList(cl storage.CategoryFilter) ([]storage.Category, error)
	CategoryIdForEdit(id int) (*storage.Category, error)
	UpdateCategory(u storage.Category) (*storage.Category, error)
}
type CoreCategory struct {
	store CategoryStore
}

func NewCoreCategory(cc CategoryStore) *CoreCategory {
	return &CoreCategory{
		store: cc,
	}
}
func (cc CoreCategory) CreateCategory(u storage.Category) (*storage.Category, error) {
	ru, err := cc.store.CreateCategory(u)
	if err != nil {
		return nil, err
	}
	if ru == nil {
		return nil, fmt.Errorf("unable to create")
	}
	return ru, nil
}
func (cc CoreCategory) DeleteCategory(id int32) error {
	err := cc.store.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
func (cc CoreCategory) CategoryList(cl storage.CategoryFilter) ([]storage.Category, error) {
	categorylist, err := cc.store.CategoryList(cl)
	if err != nil {
		return nil, err
	}
	if categorylist == nil {
		return nil, fmt.Errorf("unable to show category list")
	}

	return categorylist, nil
}
func (cc CoreCategory) CategoryIdForEdit(id int) (*storage.Category, error) {
	category, err := cc.store.CategoryIdForEdit(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, err
	}
	return category, nil
}
func (cc CoreCategory) UpdateCategory(u storage.Category) (*storage.Category, error) {
	category, err := cc.store.UpdateCategory(u)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, err
	}
	return category, nil
}
