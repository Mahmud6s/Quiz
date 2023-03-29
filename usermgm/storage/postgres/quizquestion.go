package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const insertQuizQuestionQQ = `
INSERT INTO quiz_questions (
	quiz_id,
	question_id
	
) 
VALUES ( 
	:quiz_id,
	:question_id

)RETURNING *`

func (s PostgresStorage) CreateQuizQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error) {
	stmt, err := s.DB.PrepareNamed(insertQuizQuestionQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create Quiz_question")
		return nil, fmt.Errorf("Unable to Create Quiz question")
	}
	return &u, nil
}

const EditQuizQuestionQQ = `SELECT * FROM quiz_questions WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) QuizQuestionIdForEdit(id int) (*storage.QuizQuestion, error) {
	var u storage.QuizQuestion
	if err := s.DB.Get(&u, EditQuizQuestionQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const QuizQuestionDeleteQQ = `DELETE FROM quiz_questions WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteQuizQuestion(id int32) error {
	res, err := s.DB.Exec(QuizQuestionDeleteQQ, id)
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
		return fmt.Errorf("unable to delete quiz_questions")
	}

	return nil
}

const QuizQuestionListQQ = `SELECT quiz_questions.id, questions.title, categorys.category_name, quiz.quiz_title, quiz.quiz_time 
FROM quiz_questions 
LEFT JOIN questions ON quiz_questions.question_id = questions.id 
LEFT JOIN quiz ON quiz_questions.quiz_id = quiz.id 
LEFT JOIN categorys ON questions.category_id = categorys.id 
WHERE (questions.title ILIKE '%%' || $1 || '%%' OR categorys.category_name ILIKE '%%' || $1 || '%%' 
OR quiz.quiz_title ILIKE '%%' || $1 || '%%') `

func (s PostgresStorage) QuizQuestionList(ql storage.QuizQuestionFilter) ([]storage.QzQuestion, error) {
	var List []storage.QzQuestion
	if err := s.DB.Select(&List, QuizQuestionListQQ, ql.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}

const UpdateQuizQuiestionQQ = `
	UPDATE quiz_questions SET
	quiz_id =:quiz_id,
	question_id=:question_id
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateQuiQuestion(u storage.QuizQuestion) (*storage.QuizQuestion, error) {

	stmt, err := s.DB.PrepareNamed(UpdateQuizQuiestionQQ)
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
