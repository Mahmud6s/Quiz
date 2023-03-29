package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	tests := []struct {
		name    string
		in      storage.Category
		want    *storage.Category
		wantErr bool
	}{
		{
			name: "CREATE_CATEGORY_SUCESS",
			in: storage.Category{
				ID:           1,
				CategoryName: "international",
			},
			want: &storage.Category{
				ID:           1,
				CategoryName: "international",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateCategory() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := storage.Category{
		CategoryName: "Computer",
	}
	user, err := s.CreateCategory(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(user.ID)
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_USER_SUCESS",
			in:   id,
		},
		{
			name:    "DELETE_USER_FAILD",
			in:      id,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateCategory(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := storage.Category{
		CategoryName: "sgdj",
	}
	category, err := s.CreateCategory(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := category.ID
	tests := []struct {
		name string
		in   storage.Category
		// want    *storage.Category
		wantErr bool
	}{
		{
			name: "UPDATE_CATEGORY_SUCESS",
			in: storage.Category{
				ID:           id,
				CategoryName: "sgdjUpdate",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.in.ID = category.ID
			_, err := s.UpdateCategory(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateCategory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestCategoryList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	categorys := []storage.Category{
		{
			CategoryName: "food",
		},
		{
			CategoryName: "fruits",
		},
		{
			CategoryName: "mango",
		},
	}
	for _, category := range categorys {
		_, err := s.CreateCategory(category)
		if err != nil {
			t.Fatalf("unable to create category for list category testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.CategoryFilter
		want    []storage.Category
		wantErr bool
	}{
		{
			name: "LIST_ALL_CATEGORY_SUCCESS",
			in:   storage.CategoryFilter{},
			want: []storage.Category{
				{
					CategoryName: "food",
				},
				{
					CategoryName: "fruits",
				},
				{
					CategoryName: "mango",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CategoryList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CategoryList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Category{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CategoryList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
