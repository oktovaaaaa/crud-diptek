package routes

import (
	"net/http"
	"crud-diptek/controller"
	"database/sql"	
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController()) // memanggil return dari controller
	server.HandleFunc("/students", controller.NewIndexStudent(db))
	server.HandleFunc("/students/create", controller.NewCreateStudentController(db))
	server.HandleFunc("/students/update", controller.NewUpdateStudentController(db))
	server.HandleFunc("/students/delete", controller.NewDeleteStudentController(db))

}