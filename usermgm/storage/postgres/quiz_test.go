package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateQuiz(t *testing.T) {
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
		in      storage.Quiz
		want    *storage.Quiz
		wantErr bool
	}{
		{
			name: "CREATE_QUESTION_SUCESS",
			in: storage.Quiz{
				ID:         0,
				CategoryID: int(id),
				QuizTitle:  "math1.1",
				QuizTime:   10,
			},
			want: &storage.Quiz{
				ID:         0,
				CategoryID: int(id),
				QuizTitle:  "math1.1",
				QuizTime:   10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Quiz{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuiz() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteQuiz(t *testing.T) {
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
	newQuiz := storage.Quiz{
		ID:         0,
		CategoryID: category.ID,
		QuizTitle:  "math1.1",
		QuizTime:   10,
	}
	quiz, err := s.CreateQuiz(newQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(quiz.ID)
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_QUIZ_SUCESS",
			in:   id,
		},
		{
			name:    "DELETE_QUIZ_FAILD",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateQuiz(t *testing.T) {
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
	newQuiz := storage.Quiz{
		ID:         0,
		CategoryID: category.ID,
		QuizTitle:  "math1.1",
		QuizTime:   10,
	}
	quiz, err := s.CreateQuiz(newQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	id := int32(quiz.ID)
	tests := []struct {
		name string
		in   storage.Quiz
		// want    *storage.Quiz
		wantErr bool
	}{
		{
			name: "UPDATE_QUIZ_SUCESS",
			in: storage.Quiz{
				ID:         int(id),
				CategoryID: category.ID,
				QuizTitle:  "math2.2",
				QuizTime:   11,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuizList(t *testing.T) {
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

	quiz := []storage.Quiz{
		{
			ID:         0,
			CategoryID: category.ID,
			QuizTitle:  "math1.1",
			QuizTime:   10,
		},
		{
			ID:         0,
			CategoryID: categoryTwo.ID,
			QuizTitle:  "sports1.1",
			QuizTime:   10,
		},
	}
	for _, user := range quiz {
		_, err := s.CreateQuiz(user)
		if err != nil {
			t.Fatalf("unable to create CreateQuiz for list CreateQuiz testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.QuizFilter
		want    []storage.QuizList
		wantErr bool
	}{
		{
			name: "LIST_ALL_QUESTIONS_SUCCESS",
			in:   storage.QuizFilter{},
			want: []storage.QuizList{
				{

					QuizTitle:    "math1.1",
					CategoryName: category.CategoryName,
				},
				{

					QuizTitle:    "sports1.1",
					CategoryName: categoryTwo.CategoryName,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.QuizList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.QuizList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Quiz{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.QuizList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
