package main

import (
	"net/http"

	"github.com/dimuska139/rawg-sdk-go"
	"github.com/gorilla/mux"

	"fmt"
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// lines 17 through 155 were functions created to make random users and reviews
func generateUsername() string {
	// Generate a random username
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateEmail() string {
	// Generate a random email
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var domainRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 8)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return fmt.Sprintf("%s@%s.com", string(b), string(domainRunes[rand.Intn(len(domainRunes))]))
}

func generatePassword() string {
	// Generate a random password
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var digitRunes = []rune("0123456789")
	b := make([]rune, 12)
	for i := range b {
		if i < 6 {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		} else {
			b[i] = digitRunes[rand.Intn(len(digitRunes))]
		}
	}
	return string(b)
}

func generateUsers(count int) []User {
	// Generate an array of users with unique usernames, emails, and passwords
	var users []User
	usernames := make(map[string]bool)
	emails := make(map[string]bool)

	for i := 0; i < count; i++ {
		var username string
		var email string
		for {
			username = generateUsername()
			if !usernames[username] {
				usernames[username] = true
				break
			}
		}
		for {
			email = generateEmail()
			if !emails[email] {
				emails[email] = true
				break
			}
		}
		password := generatePassword()
		users = append(users, User{Username: username, Email: email, Password: password})
	}

	return users
}

func generateReview(gameName string, username string) Review {
	// Generate a random review for the given game
	var descriptions = []string{
		"Great game, highly recommend it!",
		"Decent game, but could use some improvements.",
		"Terrible game, would not recommend it.",
		"The best game I've ever played!",
		"Average game, nothing special.",
	}

	var PS = []string{
		"DROPPED",
		"PLAYING",
		"COMPLETED",
		"ON HOLD",
	}

	return Review{GameName: gameName, Rating: float32(rand.Intn(5) + 1), Description: descriptions[rand.Intn(len(descriptions))], Username: username, PlayStatus: PS[rand.Intn(len(PS))]}
}

func generateReviews() []Review {
	// Generate an array of 10,000 unique game reviews
	var reviews []Review
	gameNames := []string{"Overwatch", "Valorant", "Destiny", "Wizard101", "Minecraft", "Xenoblade", "Pokemon"}
	gameReviews := make(map[string]map[float32]bool)
	db, err := gorm.Open(sqlite.Open("currentUsers.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	//Migrate the format of the user struct to Gorm's database
	db.AutoMigrate(&User{})
	users := generateUsers(200)
	db.Create(users)

	for _, gameName := range gameNames {
		gameReviews[gameName] = make(map[float32]bool)
	}

	for len(reviews) < 10000 {
		gameName := gameNames[rand.Intn(len(gameNames))]
		review := generateReview(gameName, users[rand.Intn(len(users))].Username)
		if !gameReviews[gameName][review.Rating] {
			gameReviews[gameName][review.Rating] = true
			reviews = append(reviews, review)
			db, err := gorm.Open(sqlite.Open("Reviews.db"), &gorm.Config{})
			if err != nil {
				panic("failed to connect to database")
			}

			//Migrate the format of the user struct to Gorm's database
			db.AutoMigrate(&Review{})
			var temp Review

			hasReview := db.Where("username = ?", review.Username, "gamename = ?", review.GameName).First(&temp).Error
			if hasReview == nil { // if review already exists, overwrite it
				UserGameRankings(&temp, false)
				temp.Rating = review.Rating
				temp.Description = review.Description
				temp.PlayStatus = review.PlayStatus
				db.Save(&temp)
				UserGameRankings(&review, true)
			} else { // else create new review
				db.Create(&Review{GameName: review.GameName, Rating: review.Rating, Description: review.Description, Username: review.Username, PlayStatus: review.PlayStatus})
				UserGameRankings(&review, true)
			}
		}
	}

	return reviews
}

// Main function -> the main point of entry
func main() {
	//generateReviews() //<- this line was used to generate random reviews and users
	
	//Creates a rounter
	router := mux.NewRouter()

	//Keeps track of who is currently signed in
	var currentlyActiveUser string = ""

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
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			SignUp(w, r)
		}
	}).Methods("POST", "OPTIONS")

	//Sign user in and ensure that the user exists
	router.HandleFunc("/sign-in", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			SignIn(w, r, &currentlyActiveUser)
		}
	}).Methods("POST", "OPTIONS", "PUT")

	//Writes a review to the review database and edits it if the desired user/game combo exists already
	router.HandleFunc("/writeareview", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			enableCors(&w)
		} else {
			WriteAReview(w, r, &currentlyActiveUser)
		}
	}).Methods("POST", "OPTIONS", "PUT")

	//Returns a list of all the reviews a specific user has made
	router.HandleFunc("/getreview", func(w http.ResponseWriter, r *http.Request) {
		GetReviews(w, r, &currentlyActiveUser)
	}).Methods("GET")

	//Returns the 4 most recent games added to the database {CALLS RECENTGAMES}
	router.HandleFunc("/recent", func(w http.ResponseWriter, r *http.Request) {
		RecentGames(w, r, client)
	}).Methods("GET")

	router.HandleFunc("/topgames", func(w http.ResponseWriter, r *http.Request) {
		TopGames(w, r, client)
	}).Methods("GET")

	router.HandleFunc("/upcominggames", func(w http.ResponseWriter, r *http.Request) {
		UpcomingGames(w, r, client)
	}).Methods("GET")

	//Start and listen for requests
	http.ListenAndServe(":8080", router)
}
