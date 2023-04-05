package main

import (
	"net/http"
	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
	
)

// Main function -> the main point of entry
func main() {

	//Creates a rounter
	router := mux.NewRouter()

	//Keeps track of who is currently signed in
	var currentlyActiveUser string = ""

	//Create RAWG SDK config and client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	//Functions that handles the url's sent from the backend:

	//PlaceHolder for a neutral handler
	router.HandleFunc("/", Hello).Methods("GET")
	http.Handle("/", router)

	//Takes in a game from the front end that is requested, and return the requested game {CALLS GAME}
	router.HandleFunc("/specific-game/", func(w http.ResponseWriter, r *http.Request) {
		Game(w, r, client)
	}).Methods("GET")

	//Returns a json of all games in the database {CALLS ALLGAMES}
	router.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) { //games?page=#&pageSize=#
		Games(w, r, client)
	}).Methods("GET")

	//Creates a user and adds it to the database {CALLS SIGNUP}
	router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			SignUp(w, r)
		}
	}).Methods("POST", "OPTIONS")

	//Sign user in and ensure that the user exists
	router.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			SignIn(w, r, &currentlyActiveUser)
		}
	}).Methods("POST", "OPTIONS", "PUT")

	//Writes a review to the review database and edits it if the desired user/game combo exists already
	router.HandleFunc("/writeareview", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			WriteAReview(w, r, &currentlyActiveUser)
		}
	}).Methods("POST", "OPTIONS", "PUT")

	//Returns a list of all the reviews a specific user has made
	router.HandleFunc("/getreview", func(w http.ResponseWriter, r *http.Request) {
		GetReviews(w, r, &currentlyActiveUser)
	}).Methods("GET")

	//Returns the 4 most recent games added to the database {CALLS RECENTGAMES}
	router.HandleFunc("/recent", func(w http.ResponseWriter, r *http.Request) {
		RecentGames(w, r, client)
	}).Methods("GET")

	router.HandleFunc("/topgames", func(w http.ResponseWriter, r *http.Request) {
		TopGames(w, r, client)
	}).Methods("GET")

	//Start and listen for requests
	http.ListenAndServe(":8080", router)
}

