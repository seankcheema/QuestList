package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	//Creates a rounter
	router := mux.NewRouter()

	//specify endpoints, handlers, and HTTP methods
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/persons", Persons).Methods("GET")
	http.Handle("/", router)

	//Start and listen for requests
	http.ListenAndServe(":8080", router)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Update response writer
	fmt.Fprintf(w, "API is up and working")
}

func Persons(w http.ResponseWriter, r *http.Request) {
	//Declare response varaiable
	var response Response

	//Retrieve person details
	persons := prepareResponse()

	//assign person details to response
	response.Persons = persons

	//Update content type
	w.Header().Set("Content-Type", "applications/json")

	//specify HTTP status code
	w.WriteHeader(http.StatusOK)

	//Convert struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	//Update response
	w.Write(jsonResponse)
}

func prepareResponse() []Person {
	var persons []Person

	var person Person
	person.Id = 1
	person.FirstName = "Issac"
	person.LastName = "Newton"
	persons = append(persons, person)

	person.Id = 2
	person.FirstName = "Albert"
	person.LastName = "Einstein"
	persons = append(persons, person)

	return persons
}
