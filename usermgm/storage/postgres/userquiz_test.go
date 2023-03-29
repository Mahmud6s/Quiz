package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateUserQuiz(t *testing.T) {
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
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
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
	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	// id := int32(user.ID)

	tests := []struct {
		name    string
		in      storage.UserQuiz
		want    *storage.UserQuiz
		wantErr bool
	}{
		{
			name: "CREATE_USER_QUIZ_SUCESS",
			in: storage.UserQuiz{
				ID:             0,
				UsersID:        user.ID,
				QuizQuestionID: quizquestion.ID,
			},
			want: &storage.UserQuiz{
				ID:             0,
				UsersID:        user.ID,
				QuizQuestionID: quizquestion.ID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateUserQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateUserQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.UserQuiz{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateUserQuiz() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteUserQuiz(t *testing.T) {
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
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
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
	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	newUserQuiz := storage.UserQuiz{
		ID:             0,
		UsersID:        user.ID,
		QuizQuestionID: quizquestion.ID,
	}
	neWquizquestion, err := s.CreateUserQuiz(newUserQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_USER_QUIZ_SUCESS",
			in:   int32(neWquizquestion.ID),
		},
		{
			name:    "DELETE_USER_QUIZ_FAILD",
			in:      int32(neWquizquestion.ID),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUserQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUserQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUpdateUserQuiz(t *testing.T) {
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
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
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
	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	newUserQuiz := storage.UserQuiz{
		ID:             0,
		UsersID:        user.ID,
		QuizQuestionID: quizquestion.ID,
	}
	neWquizquestion, err := s.CreateUserQuiz(newUserQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.UserQuiz
		wantErr bool
	}{
		{
			name: "UPDATE_USER_QUIZ_SUCESS",
			in: storage.UserQuiz{
				ID:             neWquizquestion.ID,
				UsersID:        user.ID,
				QuizQuestionID: quizquestion.ID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.UpdateUserQuiz(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUserQuiz() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserQuizList(t *testing.T) {
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
	categorytwo, err := s.CreateCategory(createCategory_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateCategory() error = %v", err)
	}
	Quiz_One := storage.Quiz{
		ID:         0,
		CategoryID: categoryOne.ID,
		QuizTitle:  "math1.1",
		QuizTime:   10,
	}
	quizOne, err := s.CreateQuiz(Quiz_One)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
	}
	Quiz_Two := storage.Quiz{
		ID:         0,
		CategoryID: categorytwo.ID,
		QuizTitle:  "math1.1",
		QuizTime:   10,
	}
	quizTwo, err := s.CreateQuiz(Quiz_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
	}
	// Quizid := int32(quiz.ID)
	Question_one := storage.Question{
		ID:         0,
		CategoryID: categoryOne.ID,
		Title:      "Bangladesh Fottball team",
	}
	questionOne, err := s.CreateQuestion(Question_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}
	Question_two := storage.Question{
		ID:         0,
		CategoryID: categorytwo.ID,
		Title:      "Bangladesh Cricket team",
	}
	questionTwo, err := s.CreateQuestion(Question_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}
	// Questionid := int32(question.ID)
	QuizQuestion_one := storage.QuizQuestion{
		ID:         0,
		QuizID:     quizOne.ID,
		QuestionID: questionOne.ID,
	}
	quizquestionOne, err := s.CreateQuizQuestion(QuizQuestion_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuizQuestion() error = %v", err)
	}
	QuizQuestion_Two := storage.QuizQuestion{
		ID:         0,
		QuizID:     quizTwo.ID,
		QuestionID: questionTwo.ID,
	}
	quizquestionTwo, err := s.CreateQuizQuestion(QuizQuestion_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuizQuestion() error = %v", err)
	}
	User_one := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	userOne, err := s.RegisterUser(User_one)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	User_Two := storage.User{
		FirstName: "rahiml",
		LastName:  "khanl",
		Email:     "rahidm@example.com",
		Username:  "radhim",
		Password:  "161252",
	}
	userTwo, err := s.RegisterUser(User_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	UserQuiz := []storage.UserQuiz{
		{
			ID:             0,
			UsersID:        userOne.ID,
			QuizQuestionID: quizquestionOne.ID,
		},
		{
			ID:             0,
			UsersID:        userTwo.ID,
			QuizQuestionID: quizquestionTwo.ID,
		},
	}
	for _, userQuiz := range UserQuiz {
		_, err := s.CreateUserQuiz(userQuiz)
		if err != nil {
			t.Fatalf("unable to create CreateUserQuiz for list CreateQuizQuestion testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.UserQuizFilter
		want    []storage.UserQuizList
		wantErr bool
	}{
		{
			name: "LIST_ALL_USER_QUIZ_SUCCESS",
			in:   storage.UserQuizFilter{},
			want: []storage.UserQuizList{
				{
					FirstName: userOne.FirstName,
					LastName:  userOne.LastName,
					Username:  userOne.Username,
					Title:     questionOne.Title,
					QuizTitle: quizOne.QuizTitle,
				},
				{
					FirstName: userTwo.FirstName,
					LastName:  userTwo.LastName,
					Username:  userTwo.Username,
					Title:     questionTwo.Title,
					QuizTitle: quizTwo.QuizTitle,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UserQuizList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UserQuizList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.UserQuizList{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UserQuizList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
