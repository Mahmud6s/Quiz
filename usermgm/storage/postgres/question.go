package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const insertQuestionQQ = `
INSERT INTO questions (
	category_id, 
	title
) 
VALUES ( 
	:category_id,
	:title

)RETURNING *`

func (s PostgresStorage) CreateQuestion(u storage.Question) (*storage.Question, error) {
	stmt, err := s.DB.PrepareNamed(insertQuestionQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create question")
		return nil, fmt.Errorf("Unable to Create question")
	}
	return &u, nil
}

const QuestionDeleteQQ = `DELETE FROM questions WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteQuestion(id int32) error {
	res, err := s.DB.Exec(QuestionDeleteQQ, id)
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
		return fmt.Errorf("unable to delete question")
	}

	return nil
}

const EditQuestionQQ = `SELECT * FROM questions WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) QuestionIdForEdit(id int) (*storage.Question, error) {
	var u storage.Question
	if err := s.DB.Get(&u, EditQuestionQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateQuestionQQ = `
	UPDATE questions SET
	category_id =:category_id,
	title=:title
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateQuestion(u storage.Question) (*storage.Question, error) {

	stmt, err := s.DB.PrepareNamed(UpdateQuestionQQ)
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

const QuestionListQQ = `SELECT categorys.category_name, questions.title
FROM questions 
LEFT JOIN categorys ON questions.category_id = categorys.id 
WHERE (categorys.category_name ILIKE '%%' || $1 || '%%' OR questions.title ILIKE '%%' || $1 || '%%' )`

func (s PostgresStorage) QuestionList(qs storage.QuestionFilter) ([]storage.QuestionList, error) {

	var List []storage.QuestionList

	if err := s.DB.Select(&List, QuestionListQQ, qs.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}
