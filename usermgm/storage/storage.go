package storage

import (
	"database/sql"
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	ID        int          `json:"id" form:"-" db:"id"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Email     string       `json:"email" db:"email"`
	Username  string       `json:"username" db:"username"`
	Password  string       `json:"password" db:"password"`
	IsAdmin   bool         `json:"is_admin" db:"is_admin"`
	IsActive  bool         `json:"is_active" db:"is_active"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.When(u.ID == 0).Error("The username field is required."),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("username cannot contain spaces"),
			validation.Required, validation.Length(6, 20),
		),
		validation.Field(&u.Email,
			validation.Required.When(u.ID == 0).Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.When(u.ID == 0).Error("The password field is required."),
			validation.Length(6, 12).Error("filed is 6 to 12 number"),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("password cannot contain spaces"),
		),
	)
}

type Login struct {
	Username string
	Password string
}

func (l Login) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Username,
			validation.Required.Error("The username field is required."),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("username cannot contain spaces"),
			validation.Required, validation.Length(6, 20),
		),
		validation.Field(&l.Password,
			validation.Required.Error("The password field is required."),
			validation.Length(6, 12).Error("filed is 6 to 12 numbers"),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("password cannot contain spaces"),
		),
	)
}

// ..
type Category struct {
	ID           int          `json:"id" form:"-" db:"id"`
	CategoryName string       `json:"category_name" db:"category_name"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (ca Category) Validate() error {
	return validation.ValidateStruct(&ca,
		validation.Field(&ca.CategoryName,
			validation.Required.Error("The username field is required."),
		),
	)
}

type CategoryFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

// ..
type Question struct {
	ID         int          `json:"id" form:"-" db:"id"`
	CategoryID int          `json:"category_id" db:"category_id"`
	Title      string       `json:"title" db:"title"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (ca Question) Validate() error {
	return validation.ValidateStruct(&ca,
		validation.Field(&ca.CategoryID,
			validation.Required.Error("The ID field is required."),
		),
		validation.Field(&ca.Title,
			validation.Required.Error("The Title field is required."),
		),
	)
}

type QuestionList struct {
	ID           int          `json:"id" form:"-" db:"id"`
	CategoryName string       `json:"category_name" db:"category_name"`
	Title        string       `json:"title" db:"title"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type QuestionFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

// ..
type Option struct {
	ID         int          `json:"id" form:"-" db:"id"`
	QuestionID int          `json:"question_id" db:"question_id"`
	OptionName string       `json:"option_name" db:"option_name"`
	IsCorrect  bool         `json:"is_correct" db:"is_correct"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (op Option) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.QuestionID,
			validation.Required.Error("The ID field is required."),
		),
		validation.Field(&op.OptionName,
			validation.Required.Error("The OptionName field is required."),
		),
	)
}

type OptionsFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

// ..
type Quiz struct {
	ID         int          `json:"id" form:"-" db:"id"`
	CategoryID int          `json:"category_id" db:"category_id"`
	QuizTitle  string       `json:"quiz_title" db:"quiz_title"`
	QuizTime   int          `json:"quiz_time" db:"quiz_time"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

func (qz Quiz) Validate() error {
	return validation.ValidateStruct(&qz,
		validation.Field(&qz.QuizTitle,
			validation.Required.Error("The  field is required."),
		),
		validation.Field(&qz.QuizTime,
			validation.Required.Error("The QuizTime field is required."),
		),
	)
}

type QuizList struct {
	ID           int          `json:"id" form:"-" db:"id"`
	QuizTitle    string       `json:"quiz_title" db:"quiz_title"`
	CategoryName string       `json:"category_name" db:"category_name"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type QuizFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

// ..
type QuizQuestion struct {
	ID         int          `json:"id" form:"-" db:"id"`
	QuizID     int          `json:"quiz_id" db:"quiz_id"`
	QuestionID int          `json:"question_id" db:"question_id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type QzQuestion struct {
	ID           int            `json:"id" form:"-" db:"id"`
	Title        string         `json:"title" db:"title"`
	CategoryName sql.NullString `json:"category_name" db:"category_name"`
	QuizTitle    string         `json:"quiz_title" db:"quiz_title"`
	QuizTime     int            `json:"quiz_time" db:"quiz_time"`
	CreatedAt    time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" db:"updated_at"`
	DeletedAt    sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}
type QuizQuestionFilter struct {
	SearchTerm string
	Sortby     int
	Offset     int
	Limit      int
}

// ..

type UserQuiz struct {
	ID             int          `json:"id" form:"-" db:"id"`
	UsersID        int          `json:"users_id" db:"users_id"`
	QuizQuestionID int          `json:"quizquestion_id" db:"quizquestion_id"`
	CreatedAt      time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt      sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type UserQuizList struct {
	ID        int          `json:"id" form:"-" db:"id"`
	FirstName string       `json:"first_name" db:"first_name"`
	LastName  string       `json:"last_name" db:"last_name"`
	Username  string       `json:"username" db:"username"`
	Title     string       `json:"title" db:"title"`
	QuizTitle string       `json:"quiz_title" db:"quiz_title"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type UserQuizFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}

// ..answer
type Answer struct {
	ID         int          `json:"id" form:"-" db:"id"`
	UserquizID int          `json:"userquiz_id" db:"userquiz_id"`
	QuestionID int          `json:"question_id" db:"question_id"`
	OptionID   int          `json:"option_id" db:"option_id"`
	IsCorrect  bool         `json:"is_correct" db:"is_correct"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type AnserList struct {
	ID         int          `json:"id" form:"-" db:"id"`
	FirstName  string       `json:"first_name" db:"first_name"`
	LastName   string       `json:"last_name" db:"last_name"`
	Username   string       `json:"username" db:"username"`
	Title      string       `json:"title" db:"title"`
	OptionName string       `json:"option_name" db:"option_name"`
	IsCorrect  bool         `json:"is_correct" db:"is_correct"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	DeletedAt  sql.NullTime `json:"deleted_at" db:"deleted_at"`
}
type AnswerFilter struct {
	SearchTerm string
	Offset     int
	Limit      int
}
