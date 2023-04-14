package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

// Tests that we can successfully put a new user with a unique username and email into the databse
// Should fail if user already exists or if username/email are already used
// {TESTS SIGNUP function}
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

// Tests that we can find an existing user in our database and match their passwords
// Should failif the user doesn't exist
// {TESTS SIGNIN function}
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

// Tests that a new review can be written and put in the review database
// If the review exists already for a specific game by the same user, its overwrttien
// {TESTS WRITEREVIEW function}
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

// Tests that a review can be found using a username passed into the fucntion and returned
// Should fail if a review by that user doesn't exist
// {TESTS GETREVIEW function}
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

// Tests that the top rated games in our database are the one that have the highest average user rating
// May change depending on the GameRanking's database entries and this test should be modified accordingly
func TestTopGames(t *testing.T) {
	t.Parallel()

	//Create RAWG client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	// Open GameRankings db
	db, err := gorm.Open(sqlite.Open("UserGameRankings.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Desired values as of 4/13 @ 1pm
	//Highest rated game atm
	highestRatedGameName := "Forza 5"
	var highestRatedGameScore float32 = 4.5

	//Second highest rated game atm
	secondHighestRatedGameName := "Wizard101"
	var secondHighestRatedGameScore float32 = 3.5

	//Create parameters for the function and call it to get a list of the rawg jsons of the top games
	r, _ := http.NewRequest("GET", "/topgames", nil)
	w := httptest.NewRecorder()

	collectionOfTopGames := TopGames(w, r, client)
	var userRating GameRanking

	//---------------------------------------------------------------------------------------------------------------------------------------------

	//Find the supposed highest rated game in our database and store it in userRating
	gameToLookup := collectionOfTopGames[0].GameName
	fmt.Println("Game name retrived from the top of the collection: " + gameToLookup)

	db.Where("game_name = ?", gameToLookup).First(&userRating)

	//Perform first check to see if the expected values match the actual values of the highest game
	gotName := userRating.GameName
	wantName := highestRatedGameName

	fmt.Println("Got Name: " + gotName + ":::")
	fmt.Println("Want Name: " + wantName + ":::")

	if gotName != wantName {
		t.Errorf("Returned wrong game name: " + gotName)
	} else {
		fmt.Println("Successful Highest Rated Game Name Test")
	}

	gotScore := userRating.AverageRating
	wantScore := highestRatedGameScore

	if gotScore != wantScore {
		fmt.Println(gotScore)
		fmt.Println(wantScore)
		t.Errorf("Returned wrong game score")
	} else {
		fmt.Println("Successful Highest Rated Game Score Test")
	}

	//--------------------------------------------------------------------------------------------------------------------------------------------

	//Find the supposed second highest rated game in our database and store it in userRating
	var userRating2 GameRanking
	gameToLookup = collectionOfTopGames[1].GameName
	fmt.Println("Game name retrived from the top of the collection: " + gameToLookup)

	db.Where("game_name = ? ", gameToLookup).First(&userRating2)

	//Perform first check to see if the expected values match the actual values of the highest game
	gotName = userRating2.GameName
	wantName = secondHighestRatedGameName

	fmt.Println("Got Name 2: " + gotName + ":::")
	fmt.Println("Want Name 2: " + wantName + ":::")

	if gotName != wantName {
		t.Errorf("Returned wrong game name: " + gotName)
	} else {
		fmt.Println("Successful Highest Rated Game Name Test")
	}

	gotScore = userRating2.AverageRating
	wantScore = secondHighestRatedGameScore

	if gotScore != wantScore {
		t.Errorf("Returned wrong game score")
	} else {
		fmt.Println("Successful Highest Rated Game Name Test")
	}
}

func TestGetUsers(t *testing.T) {
	t.Parallel()

	//Desired user to find
	desiredUser := "UnitTest"

	//Create parameters for the function
	r, _ := http.NewRequest("GET", "/getuser", nil)
	w := httptest.NewRecorder()

	//Get the list of users that have some resemblance to the provided name
	listOfPossibleUsers := GetUsers(w, r)
	foundUser := false

	fmt.Println(listOfPossibleUsers[0].Username)

	//Check to see exact name can be found in the list
	for i := 0; i < len(listOfPossibleUsers); i++ {
		if listOfPossibleUsers[i].Username == desiredUser {
			foundUser = true
			fmt.Println("Found user that matches the exact name requested: " + listOfPossibleUsers[i].Username)
			break
		}
	}

	//If exact name can't be found, see if any name exists that includes those characters, else, user cannot be found
	if !foundUser {
		for i := 0; i < len(listOfPossibleUsers); i++ {
			if strings.Contains(listOfPossibleUsers[i].Username, desiredUser) {
				foundUser = true
				fmt.Println("Found user that includes the text \"" + desiredUser + "\":" + listOfPossibleUsers[i].Username)
			}
		}
	}

	if !foundUser {
		t.Errorf("No user with the username + \"" + desiredUser + "\"" + " exists nor were there any names that matched it")
	}
}

// Test RecentReviews() which should pull the reviews from the last month in order of most recent to least recent
func TestRecentReviews(t *testing.T) {
	t.Parallel()

	//Create parameters for the function
	r, _ := http.NewRequest("GET", "/recentreviews", nil)
	w := httptest.NewRecorder()

	//Call function and set desired return
	reviews := RecentReviews(w, r)
	desired := Review{GameName: "Forza 5", Rating: float32(4.5), Description: "CAR GO VROOM", Username: "UnitTest", PlayStatus: "DROPPED"}

	if reviews != nil { // if there are reviews, proceed
		review := reviews[0] // Get most recent review

		fmt.Println("Game Name: ", review.GameName, "\nRating: ", review.Rating, "\nDescription: ", review.Description, "\nUsername: ", review.Username, "\nPlay Status: ", review.PlayStatus)
		if desired.GameName != review.GameName || desired.Rating != float32(review.Rating) || desired.Description != review.Description || desired.Username != review.Username || desired.PlayStatus != review.PlayStatus {
			t.Errorf("Review not found successfully")
		} else {
			fmt.Println("Most recent review found in database")
		}

	} else { // else fail the test
		t.Errorf("No reviews returned from RecentReviews()")
	}

}

// Tests GetFeaturedGame() which should pull the game with the highest number of reviews in UserGameRankings.db
func TestFeaturedGame(t *testing.T) {
	t.Parallel()

	//Create RAWG client
	config := rawg.Config{
		ApiKey:   "476cd66f8e4d44eb975aad199e0d7a07", //RAWG API key
		Language: "en",                               // English
		Rps:      5,                                  // Has to stay 5 (limit)
	}

	//Setup client to talk to database
	var client *rawg.Client = rawg.NewClient(http.DefaultClient, &config)

	//Create parameters for the function
	r, _ := http.NewRequest("GET", "/featuredgame", nil)
	w := httptest.NewRecorder()

	featuredGame := GetFeaturedGame(w, r, client)
	desiredGame := "Destiny"

	if featuredGame != nil{
		fmt.Println(featuredGame.Name)
		if featuredGame.Name == desiredGame{
			fmt.Println("Successfully found featured game: ", featuredGame.Name)
		} else{
			t.Errorf("Incorrect featured game")
		}
	} else{
		t.Errorf("Did not find featured game")
	}
}
