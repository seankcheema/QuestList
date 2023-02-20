package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"encoding/json"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	// _ "gorm.io/gorm"
)

// Main function -> the main point of entry
func main() {

	//Creates a rounter
	router := mux.NewRouter()
	//Create RAWG SDK config and client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)
	users := make(map[string]*user)

	//Functions that handles the url's sent from the backend:

	//PlaceHolder for a neutral handler
	router.HandleFunc("/", Hello).Methods("GET")
	http.Handle("/", router)

	//Takes in a game from the front end that is requested, and return the requested game {CALLS GAME}
	router.HandleFunc("/specific-game/{slug}", func(w http.ResponseWriter, r *http.Request) {
		Game(w, r, client)
	}).Methods("GET")

	//Returns a json of all games in the database {CALLS ALLGAMES}
	router.HandleFunc("/allGames/{page}", func(w http.ResponseWriter, r *http.Request) {
		AllGames(w, r, client)
	}).Methods("GET")

	//Creates a user and adds it to the database {CALLS SIGNUP}
	router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		SignUp(w, r, users)
	}).Methods("GET")

	//Returns the 4 most recent games added to the database {CALLS RECENTGAMES}
	router.HandleFunc("/recent", func(w http.ResponseWriter, r *http.Request) {
		RecentGames(w, r, client)
	}).Methods("GET")

	//Start and listen for requests
	http.ListenAndServe(":8080", router)
}

// Enable the front end to access backend, enables Cross-Origin Resource Sharing because frontend and backend serve from different domains
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// User Struct
type user struct {
	username string
	password string
}

// Placeholder that handles base hanlder "/"
func Hello(w http.ResponseWriter, r *http.Request) {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	fmt.Fprint(w, "Hello, Welcome to the Temporary Back-End Home Page")
}

// Handles creation of user struct and stores in the database {W-I-P}
func SignUp(w http.ResponseWriter, r *http.Request, users map[string]*user) {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	//Updates the header to indicate successful reach of the fuction
	w.WriteHeader(http.StatusOK)

	//User map Creation
	var username string
	var password string

	fmt.Println("Input Username:")
	fmt.Scanln(&username)
	fmt.Println("Input Password:")
	fmt.Scanln(&password)
	if _, ok := users[username]; ok {
		fmt.Fprint(w, "User ", username, " already exists!")
	} else {
		users[username] = NewUser(username, password)
		fmt.Fprint(w, "User ", username, " added!")
	}
}

// Helper function to help create tge struct and storing in the database
func NewUser(username string, password string) *user {
	u := user{username: username, password: password}
	return &u
}

// Takes the handler, get the game requested, and returns json
func Game(w http.ResponseWriter, r *http.Request, client *rawg.Client) {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Recieve game name from front, using the game's slug
	params := mux.Vars(r)
	slug := params["slug"]

	//--------------------

	//Update response writer
	filter := rawg.NewGamesFilter().SetPageSize(10).SetSearch(slug)
	var games []*rawg.Game
	var num int
	var err error
	games, num, err = client.GetGames(filter)

	response, err := json.Marshal(games)
	if err != nil {
		return
	}

	w.Write(response)
	if err != nil {
		return
	}

	_ = err
	_ = num
	_ = games
}

// Takes the handler's page, and returns all games of that page (40 max)
func AllGames(w http.ResponseWriter, r *http.Request, client *rawg.Client) []*rawg.Game {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Page iterator
	params := mux.Vars(r)
	tempCurrPage := params["page"]

	//cast to int
	currPage, _ := strconv.Atoi(tempCurrPage)

	//Update response writer and request all games
	filter := rawg.NewGamesFilter().SetPage(currPage).SetPageSize(40)
	var games []*rawg.Game
	var num int
	var err error

	games, num, err = client.GetGames(filter)

	//Limit of 40 games per "page"
	response, err := json.Marshal(games)
	if err != nil {
		return nil
	}

	w.Write(response)
	if err != nil {
		return nil
	}
	

	_ = err
	_ = num
	_ = games
	return games
}

// Handles requests to get the 4 most recent games released
func RecentGames(w http.ResponseWriter, r *http.Request, client *rawg.Client) {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Create time frame
	start := time.Now()
	end := start.AddDate(0, -1, 0) //1 month ago from current time

	var specifiedTime rawg.DateRange
	specifiedTime.From = end
	specifiedTime.To = start

	//Set filer to search all games in the past month, ordered by release date {handled by RAWG itself}
	filter := rawg.NewGamesFilter().SetPageSize(4).SetOrdering("released")
	var games []*rawg.Game
	var num int
	var err error

	games, num, err = client.GetGames(filter)

	response, err := json.Marshal(games)
	if err != nil {
		return
	}

	w.Write(response)

	_ = err
	_ = num
	_ = games
}
