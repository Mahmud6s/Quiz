package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const insertQuizQQ = `
INSERT INTO quiz (
	category_id, 
	quiz_title,
	quiz_time
) 
VALUES ( 
	:category_id,
	:quiz_title,
	:quiz_time

)RETURNING *`

func (s PostgresStorage) CreateQuiz(u storage.Quiz) (*storage.Quiz, error) {
	stmt, err := s.DB.PrepareNamed(insertQuizQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create quiz")
		return nil, fmt.Errorf("Unable to Create quiz")
	}
	return &u, nil
}

const QuizDeleteQQ = `DELETE FROM quiz WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteQuiz(id int32) error {
	res, err := s.DB.Exec(QuizDeleteQQ, id)
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
		return fmt.Errorf("unable to delete quiz")
	}

	return nil
}

const EditQuizQQ = `SELECT * FROM quiz WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) QuizIdForEdit(id int) (*storage.Quiz, error) {
	var u storage.Quiz
	if err := s.DB.Get(&u, EditQuizQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateQuizQQ = `
	UPDATE quiz SET
	category_id =:category_id,
	quiz_title=:quiz_title,
	quiz_time=:quiz_time
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateQuiz(u storage.Quiz) (*storage.Quiz, error) {

	stmt, err := s.DB.PrepareNamed(UpdateQuizQQ)
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

const QuizListQQ = `SELECT categorys.category_name, quiz.quiz_title
FROM quiz
LEFT JOIN categorys
ON quiz.category_id = categorys.id
WHERE (categorys.category_name ILIKE '%%' || $1 || '%%' OR quiz.quiz_title ILIKE '%%' || $1 || '%%')`

func (s PostgresStorage) QuizList(qq storage.QuizFilter) ([]storage.QuizList, error) {

	var List []storage.QuizList

	if err := s.DB.Select(&List, QuizListQQ, qq.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}
