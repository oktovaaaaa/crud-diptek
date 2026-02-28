package controller

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewUpdateStudentController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			nim := r.URL.Query().Get("nim")
			if nim == "" {
				http.Error(w, "nim harus ada", http.StatusBadRequest)
				return
			}

			var s Student
			row := db.QueryRow("SELECT nama, nim, prodi, asrama FROM students WHERE nim = $1", nim)
			if err := row.Scan(&s.Nama, &s.Nim, &s.Prodi, &s.Asrama); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := map[string]any{
				"student": s,
			}

			if err := tmpl.Execute(w, data); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}


		if r.Method == http.MethodPost {
			nimParam := r.URL.Query().Get("nim") // nim lama (ID)
			if nimParam == "" {
				http.Error(w, "nim harus ada", http.StatusBadRequest)
				return
			}

			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			nama := r.FormValue("nama")
			nim := r.FormValue("nim")
			prodi := r.FormValue("prodi")
			asrama := r.FormValue("asrama")

			_, err := db.Exec(
				"UPDATE students SET nama=$1, nim=$2, prodi=$3, asrama=$4 WHERE nim=$5",
				nama, nim, prodi, asrama, nimParam,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/students", http.StatusSeeOther)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
