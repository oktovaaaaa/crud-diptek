package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewCreateStudentController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			name := r.FormValue("nama")
			nim := r.FormValue("nim")
			prodi := r.FormValue("prodi")
			asrama := r.FormValue("asrama")
			_, err := db.Exec("INSERT INTO students (nama, nim, prodi, asrama) VALUES ($1, $2, $3, $4)", name, nim, prodi, asrama)
		if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}	
			http.Redirect(w, r, "/students", http.StatusMovedPermanently)
		return  
		} else if r.Method == "GET" {

		}
		fp := filepath.Join("views", "create.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}


err = tmpl.Execute(w, nil)
		if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}

	}
}
