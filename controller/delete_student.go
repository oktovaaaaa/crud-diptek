package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteStudentController(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		if r.Method !=http.MethodPost{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		nim := r.URL.Query().Get("nim")
		if nim == ""{
			http.Error(w, "nim harus ada", http.StatusBadRequest)
			return 
		}
		_, err := db.Exec("DELETE FROM students WHERE nim = $1", nim)
		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/students", http.StatusSeeOther)

		}
	}