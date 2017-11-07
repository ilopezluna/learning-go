package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Hello, world.\n")
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetPeople\n")
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetPerson\n")
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("CreatePerson\n")
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DeletePerson\n")
}
