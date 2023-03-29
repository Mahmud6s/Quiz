package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory := storage.Category{
		ID:           0,
		CategoryName: "footblall",
	}
	category, err := s.CreateCategory(createCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(category.ID)
	tests := []struct {
		name    string
		in      storage.Question
		want    *storage.Question
		wantErr bool
	}{
		{
			name: "CREATE_QUESTION_SUCESS",
			in: storage.Question{
				ID:         0,
				CategoryID: int(id),
				Title:      "",
			},
			want: &storage.Question{
				ID:         0,
				CategoryID: int(id),
				Title:      "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory := storage.Category{
		ID:           0,
		CategoryName: "footblall",
	}
	category, err := s.CreateCategory(createCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	// id := int32(category.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: category.ID,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(question.ID)
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_QUESTION_SUCESS",
			in:   id,
		},
		{
			name:    "DELETE_QUESTION_FAILD",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateQuestion(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory := storage.Category{
		ID:           0,
		CategoryName: "footblall",
	}
	category, err := s.CreateCategory(createCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	// id := int32(category.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: category.ID,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := question.ID
	tests := []struct {
		name string
		in   storage.Question
		// want    *storage.Question
		wantErr bool
	}{
		{
			name: "UPDATE_QUESTION_SUCESS",
			in: storage.Question{
				ID:         id,
				CategoryID: category.ID,
				Title:      "Argentina Fottball team",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.in.ID = category.ID
			_, err := s.UpdateQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuestionList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory := storage.Category{
		CategoryName: "footblall",
	}
	category, err := s.CreateCategory(createCategory)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	createCategorytwo := storage.Category{
		CategoryName: "cricket",
	}
	categoryTwo, err := s.CreateCategory(createCategorytwo)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	question := []storage.Question{
		{
			ID:         0,
			CategoryID: category.ID,
			Title:      "Bangladesh Fottball team",
		},
		{
			ID:         0,
			CategoryID: categoryTwo.ID,
			Title:      "Bangladesh cricket team",
		},
	}
	for _, user := range question {
		_, err := s.CreateQuestion(user)
		if err != nil {
			t.Fatalf("unable to create CreateQuestion for list question testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.QuestionFilter
		want    []storage.QuestionList
		wantErr bool
	}{
		{
			name: "LIST_ALL_QUESTIONS_SUCCESS",
			in:   storage.QuestionFilter{},
			want: []storage.QuestionList{
				{
					CategoryName: category.CategoryName,
					Title:        "Bangladesh Fottball team",
				},
				{
					CategoryName: categoryTwo.CategoryName,
					Title:        "Bangladesh cricket team",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.QuestionList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.QuestionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Question{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.QuestionList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
