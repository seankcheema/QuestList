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

	router.HandleFunc("/specific-game", PrintGames).Methods("GET")
	router.HandleFunc("/allGames", PrintAllGames).Methods("GET")
	router.HandleFunc("/", Hello).Methods("GET")
	http.Handle("/", router)

	//Start and listen for requests
	http.ListenAndServe(":8080", router)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Welcome to the Temporary Back-End Home Page")
}

func PrintGames(w http.ResponseWriter, r *http.Request) {
	//Specify status code
	w.WriteHeader(http.StatusOK)

	//Create RAWG SDK config and client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	client := rawg.NewClient(http.DefaultClient, &config)

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
		// fmt.Fprint(w, "Rating: ")
		// fmt.Fprintln(w, games[i].Rating)
	}

	_ = err
	_ = num
	_ = games

}

func PrintAllGames(w http.ResponseWriter, r *http.Request) {
	//Specify status code
	w.WriteHeader(http.StatusOK)
	//Create RAWG SDK config and client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to databse
	client := rawg.NewClient(http.DefaultClient, &config)

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
