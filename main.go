package main

import (
	"log"
	"net/http"

	"github.com/Erikadarisman/go-mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	routes.InitRoutes(router)
	// log.Println("creating routes")
	// //specify endpoints
	// router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	// router.HandleFunc("/persons", Persons).Methods("GET")
	// http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}
