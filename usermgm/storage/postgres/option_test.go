package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateOption(t *testing.T) {
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
	id := int(category.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: id,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	questionID := int(question.ID)
	tests := []struct {
		name    string
		in      storage.Option
		want    *storage.Option
		wantErr bool
	}{
		{
			name: "CREATE_OPTION_SUCESS",
			in: storage.Option{
				ID:         0,
				QuestionID: questionID,
				OptionName: "option1",
				IsCorrect:  false,
			},
			want: &storage.Option{
				ID:         0,
				QuestionID: questionID,
				OptionName: "option1",
				IsCorrect:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateOption(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Option{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateOption() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteOption(t *testing.T) {
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
	id := int(category.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: id,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	questionID := int(question.ID)
	Newoption := storage.Option{
		ID:         id,
		QuestionID: questionID,
		OptionName: "optionDELETE",
		IsCorrect:  false,
	}
	option, err := s.CreateOption(Newoption)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	optionID := option.ID
	tests := []struct {
		name string
		in   int32

		wantErr bool
	}{
		{
			name: "DELETE_OPTION_SUCESS",
			in:   int32(optionID),
		},
		{
			name:    "DELETE_OPTION_FAILD",
			in:      int32(optionID),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteOption(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateOption(t *testing.T) {
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
	id := int(category.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: id,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	questionID := int(question.ID)
	Newoption := storage.Option{
		ID:         id,
		QuestionID: questionID,
		OptionName: "optionDELETE",
		IsCorrect:  false,
	}
	option, err := s.CreateOption(Newoption)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	optionID := option.ID
	tests := []struct {
		name    string
		in      storage.Option
		wantErr bool
	}{
		{
			name: "UPDATE_OPTION_SUCESS",
			in: storage.Option{
				ID:         optionID,
				QuestionID: questionID,
				OptionName: "OptionUPDATE",
				IsCorrect:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateOption(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestOptionList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory_one := storage.Category{
		CategoryName: "footblall",
	}
	categoryOne, err := s.CreateCategory(createCategory_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	createCategory_two := storage.Category{
		CategoryName: "Cricket",
	}
	categoryTwo, err := s.CreateCategory(createCategory_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}

	Question_one := storage.Question{
		CategoryID: categoryOne.ID,
		Title:      "Bangladesh Fottball team",
	}
	questionOne, err := s.CreateQuestion(Question_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	Question_two := storage.Question{
		CategoryID: categoryTwo.ID,
		Title:      "Bangladesh cricket team",
	}
	questionTwo, err := s.CreateQuestion(Question_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	option := []storage.Option{
		{
			QuestionID: questionOne.ID,
			OptionName: "optionONE",
			IsCorrect:  false,
		},
		{
			QuestionID: questionTwo.ID,
			OptionName: "optionTWO",
			IsCorrect:  false,
		},
	}
	for _, op := range option {
		_, err := s.CreateOption(op)
		if err != nil {
			t.Fatalf("unable to create CreateOption for list CreateOption testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.OptionsFilter
		want    []storage.Option
		wantErr bool
	}{
		{
			name: "LIST_OPTION_SUCESS",
			in:   storage.OptionsFilter{},
			want: []storage.Option{
				{

					QuestionID: questionOne.ID,
					OptionName: "optionONE",
					IsCorrect:  false,
				},
				{

					QuestionID: questionTwo.ID,
					OptionName: "optionTWO",
					IsCorrect:  false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.OptionList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.OptionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Option{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.OptionList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
