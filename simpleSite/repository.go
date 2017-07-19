package simpleSite

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	Db *sqlx.DB
}

func (r *Repository) createUser(u *User) (int, error) {
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

func (r *Repository) updateUser(u *User) error {
	_, err := r.Db.NamedExec("update users set username=:username, age=:age, name=:name, description=:description where id=:id", &u)

	if err != nil {
		log.Println(err)
	}
	return err
}

func (r *Repository) findUserByUsername(username string) (*User, error) {
	u := User{}

	err := r.Db.Get(&u, "select * from users where username=?", username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err)
		}
		return &User{}, err
	}

	return &u, nil
}
