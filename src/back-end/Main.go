package main

//This is a test code to ensure that back end can retrieve url handles and process them

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dimuska139/rawg-sdk-go"
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

// type Game struct{
// 	Name string
// 	Publisher string
// }

func main() {

	// RAWG SDK
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}
	client := rawg.NewClient(http.DefaultClient, &config)

	filter := rawg.NewGamesFilter()
	games, num, err = client.GetGames(filter)

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
