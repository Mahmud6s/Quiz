package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const registerCategoryQQ = `INSERT INTO categorys (
	category_name
	
) VALUES(
	:category_name
	
) RETURNING *`

func (s PostgresStorage) CreateCategory(u storage.Category) (*storage.Category, error) {
	stmt, err := s.DB.PrepareNamed(registerCategoryQQ)
	if err != nil {
		return nil, err
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("unable to insert user into db")
	}
	return &u, nil
}

const deleteCategoryQQ = `DELETE FROM categorys WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteCategory(id int32) error {
	res, err := s.DB.Exec(deleteCategoryQQ, id)
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
		return fmt.Errorf("unable to delete category")
	}

	return nil
}

const CategoryListQQ = `SELECT * FROM categorys WHERE deleted_at IS NULL AND (category_name ILIKE '%%' || $1 || '%%') order by id asc`

func (s PostgresStorage) CategoryList(cl storage.CategoryFilter) ([]storage.Category, error) {

	var CategoryList []storage.Category

	if err := s.DB.Select(&CategoryList, CategoryListQQ, cl.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return CategoryList, nil
}

const EditCategoryQQ = `SELECT * FROM categorys WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) CategoryIdForEdit(id int) (*storage.Category, error) {
	var u storage.Category
	if err := s.DB.Get(&u, EditCategoryQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateCategoryQQ = `
	UPDATE categorys SET
	category_name =:category_name
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateCategory(u storage.Category) (*storage.Category, error) {

	stmt, err := s.DB.PrepareNamed(UpdateCategoryQQ)
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
