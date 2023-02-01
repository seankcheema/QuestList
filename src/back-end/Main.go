package main

//This is a test code to ensure that back end can retrieve url handles and process them

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

	// //Create RAWG SDK config and client
	// config := rawg.Config{
	// 	ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
	// 	Language: "en",                               // English
	// 	Rps:      5,                                  // Has to stay 5 (limit)
	// }
	// client := rawg.NewClient(http.DefaultClient, &config)

	// _ = client
	//specify endpoints, handlers, and HTTP methods
	router.HandleFunc("/games", PrintGames).Methods("GET")
	http.Handle("/", router)

	//Start and listen for requests
	http.ListenAndServe(":8080", router)

	// RAWG SDK

	// filter := rawg.NewGamesFilter().SetPageSize(40)
	// var games []*rawg.Game
	// var num int
	// var err error
	// games, num, err = client.GetGames(filter)
	// limit := num
	// var j int = 1
	// for limit > 0 {
	// 	for i := 0; i < 40; i++ {
	// 		fmt.Println(games[i].Name)
	// 	}
	// 	j++
	// 	filter := rawg.NewGamesFilter().SetPage(j).SetPageSize(40)
	// 	games, num, err = client.GetGames(filter)
	// 	limit -= 40
	// }

	// _ = err
	// _ = num
	// _ = games

	// if err != nil {

	// } else {
	// 	fmt.Println(err)
	// }

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
	client := rawg.NewClient(http.DefaultClient, &config)

	//Update response writer
	filter := rawg.NewGamesFilter().SetPageSize(40).SetSearch("Overwatch 20")
	var games []*rawg.Game
	var num int
	var err error
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
