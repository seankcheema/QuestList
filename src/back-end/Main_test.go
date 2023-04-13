package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
	r, _ := http.NewRequest("GET", "specific-game", nil)
	w := httptest.NewRecorder()

	values := r.URL.Query()
	values.Add("slug", "overwatch-2")
	r.URL.RawQuery = values.Encode()

	got := Game(w, r, client)[0].Name
	want := "Overwatch 2"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Println("Successful Games Test")
	}

	//TEST 2: Destiny 2
	r, _ = http.NewRequest("GET", "specific-game", nil)
	w = httptest.NewRecorder()

	values = r.URL.Query()
	values.Set("slug", "destiny-2")
	r.URL.RawQuery = values.Encode()

	got = Game(w, r, client)[0].Name
	want = "Destiny 2"

	if got != want {
		t.Errorf("Returned wrong game name: " + got)
	} else {
		fmt.Println("Successful Games Test")
	}

	//TEST 3: Rocket League
	r, _ = http.NewRequest("GET", "specific-game", nil)
	w = httptest.NewRecorder()

	values = r.URL.Query()
	values.Set("slug", "rocket-league")
	r.URL.RawQuery = values.Encode()

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

func TestSignUp(t *testing.T) {
	t.Parallel()

	db, err := gorm.Open(sqlite.Open("currentUsers.db"), &gorm.Config{}) // open db
	if err != nil {
		panic("failed to connect to database")
	}

	r, _ := http.NewRequest("POST", "sign-up", nil)
	w := httptest.NewRecorder()

	successful := SignUp(w, r) // call sign up and add user to DB

	if successful == nil {
		t.Errorf("User already exists!")
	} else {

		var username, email, password string // desired username, email, and password

		username = "UnitTest"        // should be unique
		email = "UnitTest@gmail.com" // should be unique
		password = "PASSWORD"

		var user User

		db.Where("username = ?", username).First(&user)

		fmt.Println("Username: ", user.Username, "\nEmail: ", user.Email, "\nPassword: ", user.Password)
		if user.Username != username || user.Email != email || user.Password != password {
			t.Errorf("User not added successfully")
		} else {
			fmt.Println("User successfully added to database")
		}
	}

}

func TestSignIn(t *testing.T) {
	t.Parallel()

	db, err := gorm.Open(sqlite.Open("currentUsers.db"), &gorm.Config{}) // open db
	if err != nil {
		panic("failed to connect to database")
	}

	var username, email, password string // desired username, email, and password

	username = "UnitTest"        // should be unique
	email = "UnitTest@gmail.com" // should be unique
	password = "PASSWORD"

	r, _ := http.NewRequest("POST", "sign-in", nil)
	w := httptest.NewRecorder()

	user := SignIn(w, r, nil)

	if user == nil {
		t.Errorf("User doesn't exist!")
	} else {

		db.Where("username = ?", username).First(&user)

		fmt.Println("Username: ", user.Username, "\nEmail: ", user.Email, "\nPassword: ", user.Password)
		if user.Username != username || user.Email != email || user.Password != password {
			t.Errorf("User does not exist")
		} else {
			fmt.Println("User successfully found database")
		}
	}

}

func TestWriteReview(t *testing.T) {
	t.Parallel()

	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{}) // open db
	if err != nil {
		panic("failed to connect to database")
	}

	// Desired vars
	GameName := "Forza 5"
	Rating := 4.5
	Description := "CAR GO VROOM"
	Username := "UnitTest"
	PlayStatus := "DROPPED"

	var review Review

	r, _ := http.NewRequest("POST", "/writeareview", nil)
	w := httptest.NewRecorder()

	WriteAReview(w, r, nil)

	db.Where("username = ?", Username).First(&review)

	fmt.Println("Game Name: ", review.GameName, "\nRating: ", review.Rating, "\nDescription: ", review.Description, "\nUsername: ", review.Username, "\nPlay Status: ", review.PlayStatus)
	if GameName != review.GameName || Rating != float64(review.Rating) || Description != review.Description || Username != review.Username || PlayStatus != review.PlayStatus {
		t.Errorf("Review not added successfully")
	} else {
		fmt.Println("Review found in database")
	}

}

func TestGetReview(t *testing.T) {
	t.Parallel()

	db, err := gorm.Open(sqlite.Open("reviews.db"), &gorm.Config{}) // open db
	if err != nil {
		panic("failed to connect to database")
	}

	// Desired vars
	GameName := "Forza 5"
	Rating := 4.5
	Description := "CAR GO VROOM"
	Username := "UnitTest"
	PlayStatus := "DROPPED"

	r, _ := http.NewRequest("POST", "/getreview", nil)
	w := httptest.NewRecorder()

	reviews := GetReviews(w, r, nil)
	review := reviews[0]

	db.Where("username = ?", Username).First(&review)

	fmt.Println("Game Name: ", review.GameName, "\nRating: ", review.Rating, "\nDescription: ", review.Description, "\nUsername: ", review.Username, "\nPlay Status: ", review.PlayStatus)
	if GameName != review.GameName || Rating != float64(review.Rating) || Description != review.Description || Username != review.Username || PlayStatus != review.PlayStatus {
		t.Errorf("Review not added successfully")
	} else {
		fmt.Println("Review found in database")
	}

}
