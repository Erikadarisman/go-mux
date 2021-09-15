package routes

import (
	hello "github.com/Erikadarisman/go-mux/app/services/hello"
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {

	router.HandleFunc("/", hello.Tes)
	router.HandleFunc("/check", hello.HealthCheck)
}
