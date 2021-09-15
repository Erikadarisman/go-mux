package hello

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloServiceInterface interface {
	Tes()
}

type Hello struct {
	Msg string
}

func Tes(w http.ResponseWriter, r *http.Request) {
	log.Println("entering Tes end point")
	data, err := json.Marshal(Hello{Msg: "Hello World!!"})
	if err != nil {
		http.Error(w, "faild-convert-json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}
