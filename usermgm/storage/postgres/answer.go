package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const AnswerQQ = `
INSERT INTO answer (
	userquiz_id, 
	question_id,
	option_id,
	is_correct
) 
VALUES ( 
	:userquiz_id,
	:question_id,
	:option_id,
	:is_correct

)RETURNING *`

func (s PostgresStorage) CreateAnswer(u storage.Answer) (*storage.Answer, error) {
	//..ans
	opt, err := s.GetOptionIdForEdit(u.OptionID)
	if err != nil {
		log.Println(err)
	}
	u.IsCorrect = opt.IsCorrect
	//..
	stmt, err := s.DB.PrepareNamed(AnswerQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create answer")
		return nil, fmt.Errorf("Unable to Create answer")
	}
	return &u, nil
}

const EditAnswerQQ = `SELECT * FROM answer WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) AnswerIdForEdit(id int) (*storage.Answer, error) {
	var u storage.Answer
	if err := s.DB.Get(&u, EditAnswerQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateAnswerQQ = `
	UPDATE answer SET
	userquiz_id =:userquiz_id,
	question_id =:question_id,
	option_id =:option_id,
	is_correct =:is_correct
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateAnswer(u storage.Answer) (*storage.Answer, error) {

	stmt, err := s.DB.PrepareNamed(UpdateAnswerQQ)
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

const AnswerDeleteQQ = `DELETE FROM answer WHERE id = $1 RETURNING id`

func (s PostgresStorage) AnswerDelete(id int32) error {
	res, err := s.DB.Exec(AnswerDeleteQQ, id)
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
		return fmt.Errorf("unable to delete Answer")
	}

	return nil
}

const AnswerListQQ = `SELECT users.first_name, users.last_name, users.username, questions.title, options.option_name, answer.is_correct
FROM answer
INNER JOIN user_quiz ON user_quiz.id = answer.userquiz_id
INNER JOIN users ON users.id = user_quiz.users_id
INNER JOIN questions ON questions.id = answer.question_id
INNER JOIN options ON options.id = answer.option_id
WHERE (users.username ILIKE '%%' || $1 || '%%' )`

func (s PostgresStorage) AnswerList(as storage.AnswerFilter) ([]storage.AnserList, error) {

	var List []storage.AnserList

	if err := s.DB.Select(&List, AnswerListQQ, as.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}
