package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestCreateAnswer(t *testing.T) {
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
	neWUserQuizID, err := s.CreateUserQuiz(newUserQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	newOption := storage.Option{
		ID:         0,
		QuestionID: question.ID,
		OptionName: "newOPTION",
		IsCorrect:  false,
	}
	neWoption, err := s.CreateOption(newOption)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.Answer
		want    *storage.Answer
		wantErr bool
	}{
		{
			name: "CREATE_ANSWER_SUCESS",
			in: storage.Answer{
				ID:         0,
				UserquizID: neWUserQuizID.ID,
				QuestionID: question.ID,
				OptionID:   neWoption.ID,
				IsCorrect:  false,
			},
			want: &storage.Answer{
				ID:         0,
				UserquizID: neWUserQuizID.ID,
				QuestionID: question.ID,
				OptionID:   neWoption.ID,
				IsCorrect:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.CreateAnswer(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.CreateAnswer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.Answer{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.CreateAnswer() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateAnswer(t *testing.T) {
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
	neWUserQuizID, err := s.CreateUserQuiz(newUserQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	newOption := storage.Option{
		ID:         0,
		QuestionID: question.ID,
		OptionName: "newOPTION",
		IsCorrect:  false,
	}
	neWoption, err := s.CreateOption(newOption)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	newAnswer := storage.Answer{
		ID:         0,
		UserquizID: neWUserQuizID.ID,
		QuestionID: question.ID,
		OptionID:   neWoption.ID,
		IsCorrect:  false,
	}
	newAnswerID, err := s.CreateAnswer(newAnswer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateAnswer() error = %v", err)
	}
	tests := []struct {
		name    string
		in      storage.Answer
		want    *storage.Answer
		wantErr bool
	}{
		{
			name: "UPDATE_ANSWER_SUCESS",
			in: storage.Answer{
				ID:         newAnswerID.ID,
				UserquizID: neWUserQuizID.ID,
				QuestionID: question.ID,
				OptionID:   neWoption.ID,
				IsCorrect:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tt.in.ID = category.ID
			_, err := s.UpdateAnswer(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateAnswer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAnswerDelete(t *testing.T) {
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
	neWUserQuizID, err := s.CreateUserQuiz(newUserQuiz)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	newOption := storage.Option{
		ID:         0,
		QuestionID: question.ID,
		OptionName: "newOPTION",
		IsCorrect:  false,
	}
	neWoption, err := s.CreateOption(newOption)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	newAnswer := storage.Answer{
		ID:         0,
		UserquizID: neWUserQuizID.ID,
		QuestionID: question.ID,
		OptionID:   neWoption.ID,
		IsCorrect:  false,
	}
	newAnswerID, err := s.CreateAnswer(newAnswer)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateAnswer() error = %v", err)
	}
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_ANSWER_SUCESS",
			in:   int32(newAnswerID.ID),
		},
		{
			name:    "DELETE_ANSWER_FAILD",
			in:      int32(newAnswerID.ID),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.AnswerDelete(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.AnswerDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestAnswerList(t *testing.T) {
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
		CategoryName: "cricket",
	}
	categoryTwo, err := s.CreateCategory(createCategory_two)
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
		CategoryID: categoryTwo.ID,
		QuizTitle:  "sports1.1",
		QuizTime:   10,
	}
	quizTwo, err := s.CreateQuiz(Quiz_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuiz() error = %v", err)
	}

	Question_One := storage.Question{
		ID:         0,
		CategoryID: quizOne.ID,
		Title:      "Bangladesh Fottball team",
	}
	questionOne, err := s.CreateQuestion(Question_One)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}
	Question_Two := storage.Question{
		ID:         0,
		CategoryID: quizTwo.ID,
		Title:      "Bangladesh Fottball team",
	}
	questionTwo, err := s.CreateQuestion(Question_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateQuestion() error = %v", err)
	}

	QuizQuestion_One := storage.QuizQuestion{
		ID:         0,
		QuizID:     quizOne.ID,
		QuestionID: questionOne.ID,
	}
	quizquestionOne, err := s.CreateQuizQuestion(QuizQuestion_One)
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

	newUser_one := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	userOne, err := s.RegisterUser(newUser_one)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	newUser_two := storage.User{
		FirstName: "rahimm",
		LastName:  "khann",
		Email:     "rahiim@example.com",
		Username:  "rahimm",
		Password:  "161252",
	}
	userTwo, err := s.RegisterUser(newUser_two)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}

	UserQuiz_one := storage.UserQuiz{
		ID:             0,
		UsersID:        userOne.ID,
		QuizQuestionID: quizquestionOne.ID,
	}
	UserQuizOne, err := s.CreateUserQuiz(UserQuiz_one)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}
	UserQuiz_two := storage.UserQuiz{
		ID:             0,
		UsersID:        userTwo.ID,
		QuizQuestionID: quizquestionTwo.ID,
	}
	UserQuizTwo, err := s.CreateUserQuiz(UserQuiz_two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateUserQuiz() error = %v", err)
	}

	Option_One := storage.Option{
		ID:         0,
		QuestionID: questionOne.ID,
		OptionName: "newOPTION",
		IsCorrect:  false,
	}
	optionOne, err := s.CreateOption(Option_One)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}
	Option_Two := storage.Option{
		ID:         0,
		QuestionID: questionTwo.ID,
		OptionName: "newOPTIONtwo",
		IsCorrect:  false,
	}
	optionTwo, err := s.CreateOption(Option_Two)
	if err != nil {
		t.Fatalf("PostgresStorage.CreateOption() error = %v", err)
	}

	Answer := []storage.Answer{
		{
			ID:         0,
			UserquizID: UserQuizOne.ID,
			QuestionID: questionOne.ID,
			OptionID:   optionOne.ID,
			IsCorrect:  false,
		},
		{
			ID:         0,
			UserquizID: UserQuizTwo.ID,
			QuestionID: questionTwo.ID,
			OptionID:   optionTwo.ID,
			IsCorrect:  false,
		},
	}
	for _, answer := range Answer {
		_, err := s.CreateAnswer(answer)
		if err != nil {
			t.Fatalf("unable to create CreateAnswer for list CreateAnswer testing %v", err)
		}
	}

	tests := []struct {
		name    string
		in      storage.AnswerFilter
		want    []storage.AnserList
		wantErr bool
	}{
		{
			name: "LIST_ALL_Answer_SUCCESS",
			in:   storage.AnswerFilter{},
			want: []storage.AnserList{
				{
					ID:         0,
					FirstName:  userOne.FirstName,
					LastName:   userOne.LastName,
					Username:   userOne.Username,
					Title:      questionOne.Title,
					OptionName: optionOne.OptionName,
					IsCorrect:  false,
				},
				{
					ID:         0,
					FirstName:  userTwo.FirstName,
					LastName:   userTwo.LastName,
					Username:   userTwo.Username,
					Title:      questionTwo.Title,
					OptionName: optionTwo.OptionName,
					IsCorrect:  false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.AnswerList(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.AnswerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.AnserList{}, "ID", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.AnswerList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
