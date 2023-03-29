package category

import (
	"context"
	categorypb "quiz/gunk/v1/category"
	"quiz/usermgm/storage"
)

type CoreCategory interface {
	CreateCategory(u storage.Category) (*storage.Category, error)
	DeleteCategory(id int32) error
	CategoryList(cl storage.CategoryFilter) ([]storage.Category, error)
	CategoryIdForEdit(id int) (*storage.Category, error)
	UpdateCategory(u storage.Category) (*storage.Category, error)
}
type CategorySvc struct {
	categorypb.UnimplementedCategoryServiceServer
	core CoreCategory
}

func NewCategorySvc(cc CoreCategory) *CategorySvc {
	return &CategorySvc{
		core: cc,
	}
}
func (cc CategorySvc) Register(ctx context.Context, r *categorypb.RegisterRequest) (*categorypb.RegisterResponse, error) {
	category := storage.Category{
		ID:           0,
		CategoryName: r.GetCategoryName(),
	}
	if err := category.Validate(); err != nil {
		return nil, err //TODO:: will fix when implement this service in cms
	}
	c, err := cc.core.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return &categorypb.RegisterResponse{
		Category: &categorypb.Category{
			ID:           int32(c.ID),
			CategoryName: c.CategoryName,
		},
	}, nil
}
func (cc CategorySvc) CategoryDelete(ctx context.Context, r *categorypb.CategoryDeleteRequest) (*categorypb.CategoryDeleteResponse, error) {
	err := cc.core.DeleteCategory(r.ID)
	if err != nil {
		return nil, err
	}
	return &categorypb.CategoryDeleteResponse{}, nil
}
func (cc CategorySvc) CategoryList(ctx context.Context, r *categorypb.CategoryListRequest) (*categorypb.CategoryListResponse, error) {
	user := storage.CategoryFilter{
		SearchTerm: r.GetSearchTerm(),
	}
	u, err := cc.core.CategoryList(user)
	if err != nil {
		return nil, err
	}

	var totalCategory []*categorypb.Category
	for _, ct := range u {
		user := &categorypb.Category{
			ID:           int32(ct.ID),
			CategoryName: ct.CategoryName,
		}
		totalCategory = append(totalCategory, user)
	}

	return &categorypb.CategoryListResponse{
		CategoryFilterList: &categorypb.CategoryFilterList{
			TotalCategory: totalCategory,
			SearchTerm:    "",
		},
	}, nil
}
func (cc CategorySvc) CategoryEdit(ctx context.Context, r *categorypb.CategoryEditRequest) (*categorypb.CategoryEditResponse, error) {
	ru, err := cc.core.CategoryIdForEdit(int(r.Id))
	if err != nil {
		return nil, err
	}

	return &categorypb.CategoryEditResponse{
		Category: &categorypb.Category{
			ID:           int32(ru.ID),
			CategoryName: ru.CategoryName,
		},
	}, err
}
func (cc CategorySvc) CategoryUpdate(ctx context.Context, r *categorypb.CategoryUpdateRequest) (*categorypb.CategoryUpdateResponse, error) {
	category := storage.Category{
		ID:           int(r.GetID()),
		CategoryName: r.GetCategoryName(),
	}
	_, err := cc.core.UpdateCategory(category)
	if err != nil {
		return nil, err
	}
	return &categorypb.CategoryUpdateResponse{}, nil
}
