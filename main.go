package main

import (
	"fmt"
	"log"
	"net/http"
	"server/handlers"
	"server/models"

	"github.com/gorilla/mux"
)

func main() {

	models.MigrateEvent()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/task/{id:[0-9]+}", handlers.GetEvent).Methods("GET")
	mux.HandleFunc("/api/task/", handlers.GetAllEvents).Methods("GET")
	mux.HandleFunc("/api/task/", handlers.CreateEvent).Methods("POST")
	mux.HandleFunc("/api/task/{id:[0-9]+}", handlers.UpdateEvent).Methods("PUT")
	mux.HandleFunc("/api/task/{id:[0-9]+}", handlers.DeleteEvent).Methods("DELETE")

	fmt.Println("Servidor corriendo en localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
