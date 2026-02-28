package main

import (
	"net/http"
	"crud-diptek/database"
	"crud-diptek/routes"

)

func main() {
	db := database.InitDatabase()
	
	server := http.NewServeMux()
	
	
	
	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
	
}
ppppp