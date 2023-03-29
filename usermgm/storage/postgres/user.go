package postgres

import (
	"fmt"
	"log"
	"quiz/usermgm/storage"
)

const registerUserQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password
) VALUES(
	:first_name,
	:last_name,
	:username,
	:email,
	:password
	
) RETURNING *`

func (s PostgresStorage) RegisterUser(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(registerUserQuery)
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

const registerAdminQuery = `INSERT INTO users (
	first_name,
	last_name,
	username,
	email,
	password,
	is_active,
	is_admin
) VALUES(
	:first_name,
	:last_name,
	:username,
	:email,
	:password,
	true,
	true
) RETURNING *`

func (s PostgresStorage) RegisterAdmin(u storage.User) (*storage.User, error) {
	stmt, err := s.DB.PrepareNamed(registerAdminQuery)
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

// .....
const UserListQQ = `SELECT * FROM users WHERE deleted_at IS NULL order by id asc`

func (s PostgresStorage) UserList() ([]storage.User, error) {

	var UserList []storage.User

	if err := s.DB.Select(&UserList, UserListQQ); err != nil {
		log.Println(err)
		return nil, err
	}
	return UserList, nil
}

const deleteUserQQ = `DELETE FROM users WHERE id = $1 RETURNING id`

func (s PostgresStorage) DeleteUser(id int32) error {
	res, err := s.DB.Exec(deleteUserQQ, id)
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
		return fmt.Errorf("unable to delete user")
	}

	return nil
}

const EditUserQQ = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

func (s PostgresStorage) GetUserIdForEdit(id int) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, EditUserQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateUserQQ = `
	UPDATE users SET
	first_name =:first_name,
	last_name =:last_name,
	username =:username,
	email =:email,
	password =:password
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostgresStorage) UpdateUser(u storage.User) (*storage.User, error) {

	stmt, err := s.DB.PrepareNamed(UpdateUserQQ)
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

// ..
const getUserByUsernameQuery = `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`

func (ps PostgresStorage) GetUserByUsername(usernanme string) (*storage.User, error) {
	var user storage.User
	if err := ps.DB.Get(&user, getUserByUsernameQuery, usernanme); err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("unable to find user by username")
	}

	return &user, nil
}
