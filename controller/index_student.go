package controller

import (
	"html/template"
	"net/http"
	"path/filepath"
	"database/sql"
)

type Student struct {
	Nama string
	Nim string
	Prodi string	
	Asrama string
}

func NewIndexStudent(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT nama, nim, prodi, asrama FROM students")
				if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}
		defer rows.Close()
		
		var students []Student
		for rows.Next(){
			var s Student
		if err := rows.Scan(&s.Nama, &s.Nim, &s.Prodi, &s.Asrama); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
			students = append(students, s)	
		}
		fp := filepath.Join("views", "index.html") 
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}

		data := map[string]any{
			"students": students,
		}

err = tmpl.Execute(w, data)
		if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
		}

	}
}
