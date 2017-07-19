package main

import (
	"github.com/bronhuston/simple-site/simpleSite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(127.0.0.1:3306)/simplesite")
	defer db.Close()

	r := mux.NewRouter()

	//db.MustExec("create table users ( id int unsigned not null auto_increment, primary key (id), username varchar(50) not null, age int, description varchar(250), name varchar(100) );")

	//svc := simpleSite.SaveToFileService{}
	repository := &simpleSite.Repository{Db: db}
	svc := simpleSite.SaveToDBService{Repository: repository}

	r.HandleFunc("/view/{username}", simpleSite.ViewHandler(svc))
	r.HandleFunc("/edit/{username}", simpleSite.EditHandler(svc))
	r.HandleFunc("/save/{username}", simpleSite.SaveHandler(svc))
	r.HandleFunc("/json/{username}", simpleSite.JsonHandler(svc))
	r.HandleFunc("/", simpleSite.HomePageHandler())

	http.ListenAndServe(":8080", r)
}
