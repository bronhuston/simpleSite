package simpleSite

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"html/template"
)

func ViewHandler(svc *ServiceImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		user, err := svc.GetUser(username)

		if err != nil {
			http.Redirect(w, r, "/edit/"+username, http.StatusFound)
			return
		}

		renderTemplate(w, "view", user)
	}
}

var templates = template.Must(template.ParseFiles("tmpl/view.html", "tmpl/edit.html", "tmpl/homePage.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, user *User) {
	err := templates.ExecuteTemplate(w, tmpl+".html", user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func EditHandler(svc *ServiceImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		user, err := svc.GetUser(username)

		if err != nil {
			user = &User{Username: username}
		}

		renderTemplate(w, "edit", user)
	}
}

func HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		renderTemplate(w, "homePage", &User{})
	}
}

func JsonHandler(svc *ServiceImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]
		user, err := svc.GetUser(username)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		encoder := json.NewEncoder(w)
		encoder.Encode(user)
	}
}

func SaveHandler(svc *ServiceImpl) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		username := vars["username"]

		name := r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))
		description := r.FormValue("description")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		user := &User{Username: username, Name: name, Age: age, Description: []byte(description)}

		err = svc.Save(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view/"+username, http.StatusFound)
	}
}
