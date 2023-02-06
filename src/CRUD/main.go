package main

import (
	"log"
	route "main/Routes"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Student API
// @version 1.0
// @description This is a crud application with mysql backend

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	//Read
	r.HandleFunc("/students", route.GetStudents).Methods("GET")
	//Create
	r.HandleFunc("/students", route.AddStudents).Methods("POST")
	//Update
	r.HandleFunc("/student/{ID}", route.UpdateStudent).Methods("PUT")
	//Delete
	r.HandleFunc("/student/{ID}", route.DeleteStudent).Methods("DELETE")
	//Swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
