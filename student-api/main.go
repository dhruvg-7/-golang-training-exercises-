package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/student-api/handlers"
	"github.com/student-api/stores"
)

func Server() {
	r := mux.NewRouter()
	con := stores.OpenDbConection() //

	h := handlers.NewHandler(con)

	r.HandleFunc("/student", h.InsertStudent).Methods("POST")
	r.HandleFunc("/student", h.UpdateStudent).Methods("PUT")
	r.HandleFunc("/student/{id}", h.DeleteStudent).Methods("DELETE")
	r.HandleFunc("/student", h.GetAllStudent).Methods("GET")
	r.HandleFunc("/student/{id}", h.GetStudent).Methods("GET")

	err := http.ListenAndServe(":8080", r)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
func main() {

	Server()
}
