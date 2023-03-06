package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
)

// Test that our default handler goes to the placeholder screen {TESTS HELLO function}
func TestHello(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	Hello(w, r)
	want := "Hello, Welcome to the Temporary Back-End Home Page"

	if w.Body.String() != want {
		t.Errorf("Returned wrong string: " + w.Body.String())
	} else {
		fmt.Print("Successful Hello Test")
	}
}

// Tests that our /specific-game handler takes in a slug of a game and returns the game that most closely matches it {TESTS GAMES function}
func TestGame(t *testing.T) {
	t.Parallel()

	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	//Going to test 2 searches, Overwatch 2, Destiny 2, and Rocket Leauge
	//TEST 1: Overwatch 2
	r, _ := http.NewRequest("GET", "specific-game/overwatch-2", nil)
	w := httptest.NewRecorder()

	slug := map[string]string{
		"slug": "overwatch-2",
	}

	r = mux.SetURLVars(r, slug)

	got := Game(w, r, client)[0].Name
	want := "Overwatch 2"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Println("Successful Games Test")
	}

	//TEST 2: Destiny 2
	r, _ = http.NewRequest("GET", "specific-game/destiny-2", nil)
	w = httptest.NewRecorder()

	slug = map[string]string{
		"slug": "destiny-2",
	}

	r = mux.SetURLVars(r, slug)

	got = Game(w, r, client)[0].Name
	want = "Destiny 2"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Println("Successful Games Test")
	}

	//TEST 3: Rocket League
	r, _ = http.NewRequest("GET", "specific-game/destiny-2", nil)
	w = httptest.NewRecorder()

	slug = map[string]string{
		"slug": "rocket-league",
	}

	r = mux.SetURLVars(r, slug)

	got = Game(w, r, client)[0].Name
	want = "Rocket League"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Println("Successful Games Test")
	}
}

// Tests that we can return the first game in the list all games {TESTS ALLGAMES function}
func TestAllGames(t *testing.T) {
	t.Parallel()

	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	r, _ := http.NewRequest("GET", "/games", nil)
	w := httptest.NewRecorder()

	//Check page 1
	page := map[string]string{
		"page": "1",
	}

	r = mux.SetURLVars(r, page)

	got1 := Games(w, r, client)
	got := got1[0].Name
	want := "Grand Theft Auto V"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Print("Successful AllGames Test")
	}
}
