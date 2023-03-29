package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const createUserQuizQQ = `
INSERT INTO user_quiz (
	users_id, 
	quizquestion_id
	
) 
VALUES ( 
	:users_id,
	:quizquestion_id
	
)RETURNING *`

func (s PostgresStorage) CreateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error) {
	stmt, err := s.DB.PrepareNamed(createUserQuizQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create userQuiz")
		return nil, fmt.Errorf("Unable to Create UserQuiz")
	}
	return &u, nil
}

const EditUserQuizQQ = `SELECT * FROM user_quiz WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) UserQuizIdForEdit(id int) (*storage.UserQuiz, error) {
	var u storage.UserQuiz
	if err := s.DB.Get(&u, EditUserQuizQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UserQuizDeleteQQ = `DELETE FROM user_quiz WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteUserQuiz(id int32) error {
	res, err := s.DB.Exec(UserQuizDeleteQQ, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete UserQuiz")
	}

	return nil
}

const UpdateUserQuizQQ = `
	UPDATE user_quiz SET
	users_id =:users_id,
	quizquestion_id=:quizquestion_id
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateUserQuiz(u storage.UserQuiz) (*storage.UserQuiz, error) {

	stmt, err := s.DB.PrepareNamed(UpdateUserQuizQQ)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec(u)
	if err != nil {
		log.Fatalln(err)
	}
	Rcount, err := res.RowsAffected()
	if Rcount < 1 || err != nil {
		log.Fatalln(err)
	}
	return &u, nil

}

const UserQuizListQQ = `SELECT users.first_name, users.last_name, users.username, questions.title, quiz.quiz_title
FROM user_quiz
INNER JOIN users on users.id = user_quiz.users_id
INNER JOIN questions on questions.id = user_quiz.quizquestion_id
INNER JOIN quiz_questions on quiz_questions.id = user_quiz.quizquestion_id
INNER JOIN quiz on quiz.id = quiz_questions.quiz_id
WHERE (users.first_name ILIKE '%%' || $1 || '%%' OR 
users.last_name ILIKE '%%' || $1 || '%%' OR 
users.username ILIKE '%%' || $1 || '%%' OR 
questions.title ILIKE '%%' || $1 || '%%' OR
quiz.quiz_title ILIKE '%%' || $1 || '%%')`

func (s PostgresStorage) UserQuizList(uqq storage.UserQuizFilter) ([]storage.UserQuizList, error) {
	var List []storage.UserQuizList
	if err := s.DB.Select(&List, UserQuizListQQ, uqq.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}
