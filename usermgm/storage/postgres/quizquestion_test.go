package postgres

import (
	"database/sql"
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateQuizQuestion(t *testing.T) {
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
	Quizid := int32(quiz.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: category.ID,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	Questionid := int32(question.ID)
	tests := []struct {
		name    string
		in      storage.QuizQuestion
		want    *storage.QuizQuestion
		wantErr bool
	}{
		{
			name: "CREATE_QUIZ_QUESTION_SUCESS",
			in: storage.QuizQuestion{
				ID:         0,
				QuizID:     int(Quizid),
				QuestionID: int(Questionid),
			},
			want: &storage.QuizQuestion{
				ID:         0,
				QuizID:     int(Quizid),
				QuestionID: int(Questionid),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateQuizQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateQuizQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.QuizQuestion{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateQuizQuestion() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteQuizQuestion(t *testing.T) {
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
	// Quizid := int32(quiz.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: category.ID,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}
	// Questionid := int32(question.ID)
	newQuizQuestion := storage.QuizQuestion{
		ID:         0,
		QuizID:     quiz.ID,
		QuestionID: question.ID,
	}
	quizquestion, err := s.CreateQuizQuestion(newQuizQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuizQuestion() error = %v", err)
	}
	// quizQuestionID = int32(quizquestion.ID)
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_QUIZ_QUESTION_SUCESS",
			in:   int32(quizquestion.ID),
		},
		{
			name:    "DELETE_QUIZ_QUESTION_FAILD",
			in:      int32(quizquestion.ID),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteQuizQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteQueDeleteQuizQuestionstion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateQuiQuestion(t *testing.T) {
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
	Quizid := int32(quiz.ID)
	newQuestion := storage.Question{
		ID:         0,
		CategoryID: category.ID,
		Title:      "Bangladesh Fottball team",
	}
	question, err := s.CreateQuestion(newQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	Questionid := int32(question.ID)
	newQuizQuestion := storage.QuizQuestion{
		ID:         0,
		QuizID:     int(Quizid),
		QuestionID: int(Questionid),
	}
	quizQuestion, err := s.CreateQuizQuestion(newQuizQuestion)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuizQuestion() error = %v", err)
	}
	quizQuestionid := int32(quizQuestion.ID)
	tests := []struct {
		name    string
		in      storage.QuizQuestion
		wantErr bool
	}{
		{
			name: "UPDATE_QUIZ_QUESTION_SUCESS",
			in: storage.QuizQuestion{
				ID:         int(quizQuestionid),
				QuizID:     int(Quizid),
				QuestionID: int(Questionid),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateQuiQuestion(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateQuiQuestion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestQuizQuestionList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	createCategory := storage.Category{
		CategoryName: "footblall",
	}
	categoryOne, err := s.CreateCategory(createCategory)
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

	Quiz_One := storage.Quiz{
		CategoryID: categoryOne.ID,
		QuizTitle:  "math1.1",
		QuizTime:   10,
	}
	QuizOne, err := s.CreateQuiz(Quiz_One)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
	}

	Quiz_Two := storage.Quiz{
		CategoryID: categoryTwo.ID,
		QuizTitle:  "sports1.1",
		QuizTime:   10,
	}
	QuizTwo, err := s.CreateQuiz(Quiz_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
	}

	question_one := storage.Question{
		CategoryID: QuizOne.ID,
		Title:      "Bangladesh Fottball team",
	}
	QuestionOne, err := s.CreateQuestion(question_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	question_two := storage.Question{
		CategoryID: categoryTwo.ID,
		Title:      "Bangladesh Cricket team",
	}
	QuestionTwo, err := s.CreateQuestion(question_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	quizQuestion := []storage.QuizQuestion{
		{
			QuizID:     int(QuizOne.ID),
			QuestionID: int(QuestionOne.ID),
		},
		{
			QuizID:     int(QuizTwo.ID),
			QuestionID: int(QuestionTwo.ID),
		},
	}
	for _, quizQyuestion := range quizQuestion {
		_, err := s.CreateQuizQuestion(quizQyuestion)
		if err != nil {
			t.Fatalf("unable to create CreateQuizQuestion for list CreateQuizQuestion testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.QuizQuestionFilter
		want    []storage.QzQuestion
		wantErr bool
	}{
		{
			name: "LIST_ALL_QUIZ_QUESTIONS_SUCCESS",
			in:   storage.QuizQuestionFilter{},
			want: []storage.QzQuestion{
				{
					Title:        QuestionOne.Title,
					CategoryName: sql.NullString{String: categoryOne.CategoryName, Valid: true},
					QuizTitle:    QuizOne.QuizTitle,
					QuizTime:     QuizOne.QuizTime,
				},
				{
					Title:        QuestionTwo.Title,
					CategoryName: sql.NullString{String: categoryTwo.CategoryName, Valid: true},
					QuizTitle:    QuizTwo.QuizTitle,
					QuizTime:     QuizTwo.QuizTime,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.QuizQuestionList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.QuizQuestionList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.QzQuestion{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.QuizQuestionList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
