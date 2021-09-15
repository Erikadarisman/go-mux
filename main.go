// //file main.go
// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Selamat datang di skillplus</h1>")
// 	} else if r.URL.Path == "/aboutus" {
// 		fmt.Fprint(w, "<h1>Skillplus adalah web tutorial seputar informasi teknologi</h1>")
// 	} else {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
// 	}
// }
// func main() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", handlerFunc)
// 	r.HandleFunc("/aboutus", handlerFunc)
// 	http.ListenAndServe(":3000", r)
// }

package main

import (
	"log"
	"net/http"

	"github.com/Erikadarisman/go-mux/routes"
	"github.com/gorilla/mux"
)

type Response struct {
	Persons []Person `json:"persons"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

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
