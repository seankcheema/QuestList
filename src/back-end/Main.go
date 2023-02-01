package main

//This is a test code to ensure that back end can retrieve url handles and process them

import (
	"net/http"

	"github.com/dimuska139/rawg-sdk-go"
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

}
