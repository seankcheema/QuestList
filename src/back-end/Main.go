package main

import (
	"fmt"
	"net/http"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
)

// type Game struct{
// 	Name string
// 	Publisher string
// }

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
	client := rawg.NewClient(http.DefaultClient, &config)
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
	limit := num
	var j int = 1

	//Limit of 40 games per "page" so we iterarte through all pages
	for limit > 0 {
		for i := 0; i < 40; i++ {
			fmt.Fprint(w, "Name: ")
			fmt.Fprintln(w, games[i].Name)
		}

		j++
		filter := rawg.NewGamesFilter().SetPage(j).SetPageSize(40)
		games, num, err = client.GetGames(filter)
		limit -= 40
	}

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
