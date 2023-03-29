package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const createOptionQQ = `
INSERT INTO options (
	question_id, 
	option_name
) 
VALUES ( 
	:question_id,
	:option_name

)RETURNING *`

func (s PostgresStorage) CreateOption(u storage.Option) (*storage.Option, error) {
	stmt, err := s.DB.PrepareNamed(createOptionQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create option")
		return nil, fmt.Errorf("Unable to Create option")
	}
	return &u, nil
}

const deleteOptionQQ = `DELETE FROM options WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteOption(id int32) error {
	res, err := s.DB.Exec(deleteOptionQQ, id)
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
		return fmt.Errorf("unable to delete option")
	}

	return nil
}

const OptionListQQ = `SELECT * FROM options WHERE deleted_at IS NULL AND (option_name ILIKE '%%' || $1 || '%%' ) order by id asc`

func (s PostgresStorage) OptionList(op storage.OptionsFilter) ([]storage.Option, error) {

	var List []storage.Option

	if err := s.DB.Select(&List, OptionListQQ, op.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return List, nil
}

const EditOptionQQ = `SELECT * FROM options WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetOptionIdForEdit(id int) (*storage.Option, error) {
	var u storage.Option
	if err := s.DB.Get(&u, EditOptionQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateOptionQQ = `
	UPDATE options SET
	question_id =:question_id,
	option_name =:option_name,
	is_correct =:is_correct
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateOption(u storage.Option) (*storage.Option, error) {

	stmt, err := s.DB.PrepareNamed(UpdateOptionQQ)
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
