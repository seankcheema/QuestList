package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"encoding/json"

	"github.com/dimuska139/rawg-sdk-go"
	//"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
}

// Review Struct
type Review struct {
	gorm.Model
	Rating      float32
	Description string
	User        string
}

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

	//Functions that handles the url's sent from the backend:

	//PlaceHolder for a neutral handler
	router.HandleFunc("/", Hello).Methods("GET")
	http.Handle("/", router)

	//Takes in a game from the front end that is requested, and return the requested game {CALLS GAME}
	router.HandleFunc("/specific-game", func(w http.ResponseWriter, r *http.Request) {
		Game(w, r, client)
	}).Methods("GET")

	//Returns a json of all games in the database {CALLS ALLGAMES}
	router.HandleFunc("/games", func(w http.ResponseWriter, r *http.Request) { //games?page=#&pageSize=#
		Games(w, r, client)
	}).Methods("GET")

	//Creates a user and adds it to the database {CALLS SIGNUP}
	router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		SignUp(w, r)
	}).Methods("POST", "OPTIONS", "PUT")

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
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// Placeholder that handles base hanlder "/"
func Hello(w http.ResponseWriter, r *http.Request) {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	fmt.Fprint(w, "Hello, Welcome to the Temporary Back-End Home Page")
}

// Handles creation of user struct and stores in the database {W-I-P}
// ------------ THIS IS NOT INTENDED IMPLEMENTATION AND IS NOT TESTED ---------------------------------
func SignUp(w http.ResponseWriter, r *http.Request) *User {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	fmt.Print("IN SIGN_UP")

	//Updates the header to indicate successful reach of the fuction
	// w.WriteHeader(http.StatusOK)

	//Open the database
	db, err := gorm.Open(sqlite.Open("currentUsers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Migrate the format of the user struct to Gorm's database
	db.AutoMigrate(&User{})

	var user User

	//Recieve username and password from front, using the parameters listed in the passed in json file

	json.NewDecoder(r.Body).Decode(&user)

	fmt.Println("username: " + user.Username)

	//Check that the username doesn't already exist in the database
	// _, hasUser := db.Get(user.Username)
	hasUser := db.Where("username = ?", user.Username).First(&user).Error
	// fmt.Println(hasUser)

	if hasUser == nil {
		fmt.Println("User ", user.Username, " already exists!")
		//w.WriteHeader(http.StatusTeapot) //IDK What this status does
		return nil
	} else { //If its a new user, add the user and the information to the database
		db.Create(&User{Username: user.Username, Password: user.Password})
		fmt.Fprint(w, "User ", user.Username, " added!")

		w.WriteHeader(http.StatusCreated)
		return &user

	}

	// if err = db.Where("username = ?", user.Username).First(&user).Error; err != nil { //If it does exist, say user exists and return nil
	// 	fmt.Fprint(w, "User ", user.Username, " already exists!")
	// 	w.WriteHeader(http.StatusTeapot) //IDK What this status does
	// 	return nil

	// } else { //If its a new user, add the user and the information to the database
	// 	db.Create(&User{Username: user.Username, Password: user.Password})
	// 	fmt.Fprint(w, "User ", user.Username, " added!")

	// 	w.WriteHeader(http.StatusCreated)
	// 	return &user

	// }

}

// Takes the handler, get the game requested, and returns json
func Game(w http.ResponseWriter, r *http.Request, client *rawg.Client) []*rawg.Game {
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
		return nil
	}

	w.Write(response)
	if err != nil {
		return nil
	}

	_ = err
	_ = num
	_ = games

	fmt.Println(games[0].Name)

	return games
}

// Takes the handler's page, and returns all games of that page (40 max)
func Games(w http.ResponseWriter, r *http.Request, client *rawg.Client) []*rawg.Game {
	//Allows the doamin to be accessed by frontenf
	enableCors(&w)

	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Page iterator
	// params := mux.Vars(r)
	// tempCurrPage := params["page"]
	tempCurrPage := r.URL.Query().Get("page")
	tempPageSize := r.URL.Query().Get("pageSize")

	var err error
	//cast to int
	currPage, err := strconv.Atoi(tempCurrPage)
	if err != nil {
		currPage = 1
	}
	pageSize, err := strconv.Atoi(tempPageSize)
	if err != nil {
		pageSize = 40
	}

	//Update response writer and request all games
	filter := rawg.NewGamesFilter().SetPage(currPage).SetPageSize(pageSize)
	var games []*rawg.Game
	var num int

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
// Due to the constantly changing nature of the games that are updated per day, this cannot be concretely unit tested with set values
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
