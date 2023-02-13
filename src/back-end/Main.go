package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "gorm.io/gorm"
)

func main() {

	//Creates a rounter
	router := mux.NewRouter()
	//Create RAWG SDK config and client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	// GORM TESTING
	// db, err := gorm.Open("sqlite2", "Users.db")
	// if err != nil {
	// 	fmt.Println("Connection Failed")
	// } else {
	// 	fmt.Println("Connection Established")
	// }

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)
	users := make(map[string]*user)

	router.HandleFunc("/specific-game", func(w http.ResponseWriter, r *http.Request) {
		PrintGames(w, r, client)
	}).Methods("GET")
	router.HandleFunc("/allGames", func(w http.ResponseWriter, r *http.Request) {
		PrintAllGames(w, r, client)
	}).Methods("GET")
	router.HandleFunc("/sign-up", func(w http.ResponseWriter, r *http.Request) {
		SignUp(w, r, users)
	}).Methods("GET")
	router.HandleFunc("/", Hello).Methods("GET")
	http.Handle("/", router)

	//Start and listen for requests
	http.ListenAndServe(":8080", router)

}

func SignUp(w http.ResponseWriter, r *http.Request, users map[string]*user) {
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
		users[username] = newUser(username, password)
		fmt.Fprint(w, "User ", username, " added!")
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Welcome to the Temporary Back-End Home Page")
}

func PrintGames(w http.ResponseWriter, r *http.Request, client *rawg.Client) {
	//Specify status code
	w.WriteHeader(http.StatusOK)

	fmt.Println("Input game name:")

	var err error

	var name string

	fmt.Scanln(&name)

	//Update response writer
	filter := rawg.NewGamesFilter().SetPageSize(40).SetSearch(name)
	var games []*rawg.Game
	var num int
	games, num, err = client.GetGames(filter)
	for i := 0; i < 10; i++ {
		fmt.Fprint(w, "Name: ")
		fmt.Fprintln(w, games[i].Name)
		fmt.Fprint(w, "Rating: ")
		fmt.Fprintln(w, games[i].Rating)
	}

	_ = err
	_ = num
	_ = games

}

func PrintAllGames(w http.ResponseWriter, r *http.Request, client *rawg.Client) {
	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Update response writer and request all games
	filter := rawg.NewGamesFilter().SetPageSize(40)
	var games []*rawg.Game
	var num int
	var err error

	games, num, err = client.GetGames(filter)

	//Limit of 40 games per "page" so we iterarte through all pages

	// for i := 0; i < 40; i++ {
	fmt.Fprint(w, "Name: ")
	fmt.Fprintln(w, games[0])
	response, err := json.Marshal(games[0])
	if err != nil {
		return
	}
	//}
	w.Write(response)

	_ = err
	_ = num
	_ = games
}

type user struct {
	username string
	password string
}

func newUser(username string, password string) *user {
	u := user{username: username, password: password}
	return &u
}
