package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimuska139/rawg-sdk-go"
)

func AllGamesTest(t *testing.T) {
	t.Parallel()

	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	r, _ := http.NewRequest("GET", "allGames/1", nil)
	w := httptest.NewRecorder()


	got := AllGames(w, r, client)[0].Name
	want := "Grand Theft Auto V"

	if got != want {
		t.Errorf("Returned wrong game name: ", got)
	} else {
		fmt.Print("Successful AllGames Test")
	}
}
