package simpleSite

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository interface {
	CreateUser(u *User) (int, error)
	UpdateUser(u *User) error
	FindUserByUsername(s string) (*User, error)
	GetAddressesByUserName(s string) (*[]Address, error)
}

type RepositoryImpl struct {
	Db *sqlx.DB
}

func (r *RepositoryImpl) CreateUser(u *User) (int, error) {
	res, err := r.Db.NamedExec("Insert into users (username, name, age, description) values (:username, :name, :age, :description)", &u)
	if err != nil {
		log.Println(err)
	}

	createdId, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	log.Printf("User created with username: %s and id: %d", u.Username, createdId)

	return int(createdId), err
}

func (r *RepositoryImpl) UpdateUser(u *User) error {
	_, err := r.Db.NamedExec("update users set username=:username, age=:age, name=:name, description=:description where id=:id", &u)

	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *RepositoryImpl) FindUserByUsername(username string) (*User, error) {
	u := User{}

	err := r.Db.Get(&u, "select u.* from users u where username=?", username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		return &User{}, err
	}

	return &u, nil
}

func (r *RepositoryImpl) GetAddressesByUserName(username string) (*[]Address, error) {
	a := []Address{}
	err := r.Db.Select(&a, "select a.* from users u inner join addresses a on a.user_id = u.id where u.username = ?", username)
	if err != nil {
		log.Println(err)
		return &[]Address{}, err
	}

	return &a, nil
}
