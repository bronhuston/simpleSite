package simpleSite

import (
	"database/sql"
	"time"
)

type User struct {
	Username    string
	Name        string
	Age         int
	Description []byte
	Id          int
	Addresses   []Address
	CreatedTimestamp
	UpdatedTimestamp
}

type CreatedTimestamp struct {
	CreatedAt time.Time `db:"created_at"`
}

type UpdatedTimestamp struct {
	UpdatedAt time.Time `db:"updated_at"`
}

type Address struct {
	Id          sql.NullInt64
	Addr_line_1 sql.NullString
	Addr_line_2 sql.NullString
	City        sql.NullString
	State       sql.NullString
	Zip_5       sql.NullInt64
	User_id     sql.NullInt64
}
