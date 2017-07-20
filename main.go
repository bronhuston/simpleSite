package main

import (
	"github.com/bronhuston/simple-site/simpleSite"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"log"
	"net/http"
)

func main() {
	db := sqlx.MustConnect("mysql", "root:password@tcp(127.0.0.1:3306)/simplesite")
	defer db.Close()

	err := goose.SetDialect("mysql")
	if err != nil {
		log.Fatal("Unable to set goose dialect: ", err)
	}
	err = goose.Up(db.DB, "./db")
	if err != nil {
		log.Fatal("Goose migrations failed: ", err)
	}

	r := mux.NewRouter()

	repository := &simpleSite.Repository{Db: db}
	svc := &simpleSite.SaveToDBService{Repository: repository}

	r.HandleFunc("/view/{username}", simpleSite.ViewHandler(svc))
	r.HandleFunc("/edit/{username}", simpleSite.EditHandler(svc))
	r.HandleFunc("/save/{username}", simpleSite.SaveHandler(svc))
	r.HandleFunc("/json/{username}", simpleSite.JsonHandler(svc))
	r.HandleFunc("/", simpleSite.HomePageHandler())

	http.ListenAndServe(":8080", r)
}
