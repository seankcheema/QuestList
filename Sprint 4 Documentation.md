Javier Lopez, Richard Sorkin, Sean Cheema, and Thomas Gallego

4/14/23

Professor Dobra

**Sprint 4 Documentation**
<br/>

**_User Stories Worked On This Sprint:_**
<br />

**1) Home Page** <hr />As a **site visitor** I would like to **access a home page** for the purpose of **create a hub for navigation.**
<br />
**Implementation Details** <hr /> 
- Implemented by front and back-end
- Create a GUI to quickly access other pages
- Display new game inform

**2) Review Issue** <hr />As a site visitor I would like to review games for the purpose of keeping track of the games I have played/am playing and how I liked them
<br />
**Implementation Details** <hr />
- Implemented by front and back end
- Create GUI to be able to review games
- Create a review database with elements GameName, Rating, Description, Username, and PlayStatus
- If a user creates a review for a game they have already reviewed, it will overwrite that review
<br />

**Work Completed**<hr/>

**Front-end:**

_Work Completed_:<br/>
WIP


Cypress Testing:<br/>
WORK IN PROGRESS

_Testing Challenges_: <br/>
One issue we encountered was that the typescript compiler had chai type conflict errors when using Jasmine alongside Cypress. While the tests for both Cypress and Jasmine would still run, VSCode threw errors. This was important to fix, though, as it was difficult to leverage typescript’s errors for debugging purposes while these type conflict errors were present.
<hr />

**Backend:**

_Work Completed_:<br/>
In Sprint 4 we focused primarily on correcting and improving previous functions, organization of code, and functions to be used for the front-end’s homepage implementations. 

We began by refining the WriteAReview() function as it was returning faulty data. The way we were doing look-ups in the database when multiple parameters were needed (in this case username and game name) was incorrect as it was returning the first instance of the first parameter rather than taking both parameters into consideration. In this function a new function called UserGameRankings() was added with the purpose of storing combined metadata in a separate database of new GameRanking structs every time a new review is made to help with the performance of other functions. For example, with this implementation the entire review database would not need to be parsed every time the TopGames() or GetFeaturedGame() functions are called. Additionally, the Game() function was reconfigured to take in query parameters.

To organize our code, we deviated from the Main.go file and created a Main_functions.go file. The Main.go file only contains our handler functions and other necessary code such as client definition, gorilla mux variable definitions, and functions used to randomly generate users and reviews for testing purposes. Our Main_functions.go file on the other hand contains all previous functions plus the new ones added in this sprint.

Most of the functions added in this sprint were designed with homepage implementation in mind. From these functions, some issues that arose were in TopGames(), UpcomingGames(), and RecentReviews(). With TopGames(), we struggled with appending and finding games by name. We fixed this by using a slice and golang’s append function to prevent an out of bounds error that frequently popped up. Additionally, chose to use a quick sort algorithm to sort the games returned from the UserGameRankings database in descending order of rating starting with the highest rated game. With UpcomingGames(), we struggled in the time frame functionality as we had to look forward in time. To fix this issue, the date range begins one day after golang’s time.Now() which is the current time. Lastly, with RecentReviews(), the original functionality returned an array of reviews starting with the least recent since this is what GORM’s lookup returns by default, but the intended functionality should return one starting with the most recent. To account for this, the reverseArray() function was created and used to return the array in the correct order. Additional functions we added with minimal to no issues include the GetUser() function to be used for user lookup and RecentReviews() for homepage display.
<br/>

_Back End Unit Tests_: <br/>
We have implemented a variety of Unit Tests using the built-in Go testing. Initially, we created TestHello() to test the functionality of the Go “testing” import. To test our Main.go functions we created tests TestGame(), TestAllGames(), TestSignUp(), TestSignIn(), TestWriteReview(), and TestGetReview(). 

-TestGame():<br/>
TestGame() tests the function Game() with a specific game “slug” (slug is a concatenated version of a game name setting all letters to lowercase and replacing all spaces with hyphens). When this slug is passed, Game should return the 10 most similar games to that slug with array[0] being the most similar and array[9] being the least similar. For this test we pulled the first element of the array as it is the desired game and compared it to our intended game string. 

-TestAllGames():<br/>
Next, we used TestAllGames() to test the function AllGames(). AllGames() should return every game in the RAWG API one “page” at a time. Each page will have 40 games starting from the most popular game in the API to the least popular for all 800,000+ games. For this test, we pulled the first element of the first page which would be the most popular game in the API, in this case “Grand Theft Auto V”. 

-TestSignUp():<br/>
This tests the functionality of Sign-Up(). This tested function should add a new user to the database of users if their username and/or email have not been used already. For this test, we created an arbitary username, password, and email and inserted them into the database. To ensure they were added, we scan the database to find the first, and only occurrance of their username. To test the prevention of duplicate users, we try to add an already existing username and expect an error.

-TestSignIn():<br/>
This tests the functionality of Sign-In(). Sign-In() should check that a requested username exists in the database. If they do exisit, we compare the stored password with the requested one. If everything matches up, we return an OK status and mark the signed-in user as the active user. If something doesn't match the database, we return an error status. Our test first tries to find a user that does exist with a mathcing password. This should return an OK status. Afterwards, it tries logging in with a nonexistent user, resulting in an error.

-TestWriteReview():<br/>
This tests the function WriteAReview() with a review struct passed in containing elements such as the game being reviewed, rating out of 5, description, username of the user who reviewed the game, and a playstatus. The intended functionality is that WriteAReview() will add the review to the Review database if that user has not already reviewed that game, and modify the review if a user has already created a review for that game.

-TestGetReview():<br/>
This tests the function GetReview() with a user struct passed in containing elements such as a unique username, unique email, and a password. The intended functionality of GetReview() is that the function retrieves all instances of a specific user’s reviews from the review database and writes an error status to the header if the user has not created a review.

-TestTopGames():<br/>
This tests the functionality of TopGames(). TopGames() should return the most popular games based on QuestList user rankings, which are stored in a database and updated in real time as new reviews are made. The test compares the expected names and scores of the 2 highest rated games with the values found in the database "UserGameRankings." The expected values are found by observing the values in the databse and writing down the name and rating field as expected values. The TopGames() function is then called and the first two games from the returned array are compared with the expected name and ratings. If they match, the test passes.

-TestGetUser():<br/>
This tests the fucntionailty of GetUsers(). GetUsers() returns an array of user structs found in the "currentUsers" database that have names that are somewhat similar to the passed in user name that is being searched. The test can be changed to search different usernames, but for the sake of simplicity, the test searches for a precreated user called "UnitTest" to see if an existing user can be searched and returned. The test parses the array of returned users and checks if the exact name is found. If it is, the test is successful and says that the exact name was found. Otherwise, the test checks if a name containing the passed in username exisits and returns correct if one is found. A print statement is also made that alerts that it is not an exact match. If none of these tests work, the unit test fails. In our case, we get an exact match and pass the test.

-TestRecentReviews():<br/>
This tests the functionality of RecentReviews(). The intended functionality is to access the Reviews database and extract all reviews updated (and subsequently created) in the last month and returns an array of them in order of recency with the most recent review being at the 0th index. The test function pulls this most recent review and compares it with the expected review. If they are not the same or the review array is empty, the test throws an error, otherwise the test passes.

-TestFeaturedGame():<br/>
This tests the functionality of GetFeaturedGame(). The intended functionality is to access the UserGameRankings database and find the GameRanking object within that database with the highest number of reviews. Once this object is found, the name is used to find the corresponding rawg.Game object and returns it. The test function throws an error if the expected featured game is not the same as the one returned from the function, or if the returned value is nil, otherwise the test passes

-Testing Limitations:<br/>
It should be noted that RecentGames() could not be tested. For RecentGames(), there is no way to predict the outcome of RecentGames() as new games are being added to the API every day causing the RecentGames() output to be changed frequently. In a similar manner, the function "upcominggames" cannot be unit tested. This function returns games that are going to be released within the next month, starting from one day from the present. Since the dates change daily (and possibly hourly at times due to time zone differences) any unit test created would result in different values and thua cannot be properly unit tested. However, first hand use of the functions have shown that correct values are returned.

-External Library Documentation_:<br/>
The API we used is titled RAWG API. RAWG is a database of games storing a variety of information including but not limited to: name, image, different ratings, release date, and developers. We have used this data to organize and traverse the database to acquire desired information. In order to easily use the RAWG API, we found “RAWG SDK GO”, a client with built in functions to traverse the RAWG database. The documentation for both RAWG API and RAWG SDK GO are linked below. 

RAWG API: https://api.rawg.io/docs/
RAWG SDK GO: https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-documentation

<hr />

**QuestList Backend API Documentation**

_NOTE: The full detailed, colored version of the documentation with images to show the format of the JSON files can be found in a PDF file outside of this markup file_


**Introduction**
<br />
The backend API for the Quest List project specializes in retrieving and storing data passed into it by the front-end servers. It aims to allow the front-end to properly retrieve and store information about individual users in the form of sign up (storage of information) and sign in (information retrieval). When dealing with games, our API supports a variety of function calls that can provide the front-end developers with different game information depending on the specifications passed into our function calls.

_How is Information Shared?_
<br />
When communicating between the front and back-end, two different procedures are employed depending on whether information is being requested by the front or being sent to the front.

_How can the front-end call the backend?_
<br />
To call any of the functionality from the backend, the front must make a URL request to the server hosting the backend (in our case http://localhost:8080/). When calling these URLs, certain handlers must be specified after the final “/” to attain the desired results. The handlers and their functions are listed below.

_What data can be sent to the backend and in what form can it be sent?_
<br />
Information can be sent to the backend through json files. Json files can be created and formatted in a way that matches one of the structs outlined below. Structs are basically objects, so in order to properly have all values intended to be sent to the back be interpreted correctly, a similar object must be created in the front and packaged into a json file. All structs used in the back are listed below and the documentation for each handler will advise you about the requirements of json files. When certain parameters are required by the backend in the URL, we expect the front-end to use query parameters to add the necessary information which the backend then reads from the URL

_How is data sent from the backend to the front-end?_
<br />
Once an operation is completed,  the backend will return two forms of data. The first return is found in every operation handled by the backend: an http status. If a function call is successful and everything was retrieved, stored, or modified correctly, the backend will write to the header: “http.StatusOK.” If something went wrong, each function call will write a different http status to alert to potential errors. 
The second type of data returned is a json. Though not all operations return a json file, the ones that do will be specified and the layout of the file will be shown.  

_Where does the game content come from?_
<br />
The API we used is titled RAWG API. RAWG is a database of games storing a variety of information including but not limited to: name, image, different ratings, release date, and developers. We have used this data to organize and traverse the database to acquire desired information. In order to easily use the RAWG API, we found “RAWG SDK GO”, a client with built in functions to traverse the RAWG database. The documentation for both RAWG API and RAWG SDK GO are linked below.

RAWG API: https://api.rawg.io/docs/
RAWG SDK GO: https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-documentation

**Structs Used Throughout the Backend**

1) User Struct<br />

Purpose<br />
Used to store and communicate information about users of the website. Stores a username (string), an email address (string), and a password (string).

Format
<br />
type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}
  
2) Review Struct<br />

Purpose<br />
Used to store and communicate information about the reviews users have made about a specific game. It stores the name of the game (string), the rating of the game (float), the description of the review a user has commented (string), the username (string), and the play status of the game in question (string). 

Note: the playstatus can only be one of the following strings: PLAYING, DROPPED, COMPLETED, ON HOLD

Format
<br />
type Review struct {
    gorm.Model
    GameName    string  //Names of game being reviewed
    Rating      float32 //Rating (out of 5) of the game
    Description string  //Description of the game played
    Username    string  //Name of the account
    PlayStatus  string  //PLAYING, DROPPED, COMPLETED, ON HOLD
}

<br />

3) GameRanking Struct<br />

Purpose<br />
This is the structure for our overarching rating database for the games on our website. Every time a new game is added to our website or rated for the first time, an entry in the UserGameRanking database is made using this scheme. The GameName is the name of the video game, the AverageRating is an average calculated by dividing the overall sum of all user scores for a game by the number of reviews. NumReviews is the number of reviews a game has.

Format 
<br />
type GameRanking struct {
	gorm.Model
	GameName      string  `gorm:"uniqueIndex"` // Name of game
	AverageRating float32 // Average Rating (out of 5) of the game
	NumReviews    int     // Number of times a game has been reviewed
}
<hr />

**Handler Functions and their Functionality**

1) GET: {HelloWorld}: http://localhost:8080/

Functionality
<br />
This url serves as a test. It prints out a message to the backend page if the server is up and running. This is not meant to return nor do any meaningful computations and can be used to test if the backend servers are running 

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: N/A 

Json Returned & Format: 
<br />
N/A

<br />

2) GET: {Get a json of 10 games most similar to provided name}:
http://localhost:8080/specific-game

Functionality
<br />
This url retrieves the 10 games most similar to the name provided in the query parameters provided and sends a json of these 10 games. We use RAWG SDK’s NewGamesFilter() function along with SetPageSize(int numElements) and SetSearch(string gameName) to retrieve an array of *rawg.Game objects. SetPageSize(int numElements) takes in the parameter 10 to set the number of returned elements to 10. SetSearch(string gameName) then takes in the query parameter name as a parameter and retrieves the 10 most similar elements to that name. From there, we json.Marshal() these games and write these games to the header.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: N/A

Json Returned & Format: 
<br />
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

<br />

3) GET: {Get a json of all games in a specific page of the database}: http://localhost:8080/games

Functionality
<br />
This function gets a json of all the games in a page specified through the url query parameters. These parameters specify the page and page size which we then pass into SetPage(int page) and SetPageSize(int numElements) respectively along with RAWG SDK’s NewGamesFilter() function. The combination of these functions returns an array of size “numElements” from the page “page” in the game database. From there, we json.Marshal() these games and write these games to the header.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: N/A

Json Returned & Format: 
<br />
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

<br />

4) Post: {Create a new user and add it to the database}: http://localhost:8080/sign-up

Functionality
<br />
This function creates a user and adds it to the database if it is not already there. This function starts by opening the user database and migrating the format of the user struct to Gorm’s database. The function then pulls the body content from the front-end using json.Decode() to put this struct in a user object of type struct User. The function then validates the uniqueness of the email and username. If either the username or email exist in the database already, the user will not be added to the database and the http.StatusInternalServerError status will be returned. Otherwise, the user is added to the database and the http.StatusCreated status is returned.

Status Returned:
<br />
If Successful: http.StatusCreated
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format:
<br />
type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}

<br />

5) POST: {Sign a user in and assure their account exists}: http://localhost:8080/sign-in

Functionality
<br />
This function finds a user in the database and logs that user in if the user is found in the database. Similar to the SignUp() function, this function starts by opening the user database and migrating the format of the user struct to Gorm’s database. The function then pulls the body content from the front-end using json.Decode() to put this struct in a user object of type struct User. The function then validates that the specified user exists in the database. If it doesn’t, the http.StatusInternalServerError status is returned. However, if the username exists in the database, the password is then verified, returning the Internal Server Error status if it doesn’t match the one in the database. Lastly, if the password matches the user is logged in by storing that username in a currentlyActiveUser variable and writing the http.StatusOK to the header telling the front-end that the user was found.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: 
<br />
type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}

<br />

6) POST: {Create a new review for a user}: http://localhost:8080/writeareview

Functionality
<br />
This function creates a new review for a user and overwrites the user’s previous review if they have already created one for this game. This function starts by opening the review database and migrating the format of the review struct to Gorm’s database. The function then pulls the body content from the front-end using json.Decode() to put this struct in a review object of type struct Review. The function then checks if the user has already created a review for this game. If they have, the rating, description, and play status are changed for that specific review in the database and the http.StatusOK status is returned. If they have not created a review yet for this game, a new review object is created in the database in the Review struct format and the http.StatusCreated status is returned. It should be noted that the user does not have to add every element to the review (e.g. a user may create a review with no description).

Status Returned:
<br />
If Successful: http.StatusOK or http.StatusCreated
<br />
If Unsuccessful: N/A

Json Returned & Format: 
<br />
type Review struct {
    gorm.Model
    GameName    string  //Names of game being reviewed
    Rating      float32 //Rating (out of 5) of the game
    Description string  //Description of the game played
    Username    string  //Name of the account
    PlayStatus  string  //PLAYING, DROPPED, COMPLETED, ON HOLD
}

<br />

7) GET: {Get a json with a list of a specified user’s review structs}: http://localhost:8080/getreviews

Functionality
<br />
This function takes in a parameter of type user struct and retrieves all of the reviews created by that user. This function starts by opening the review database and migrating the format of the user struct to Gorm’s database. The function then pulls the body content from the front-end using json.Decode() to put this struct in a user object of type struct User. The function then searches in the database for all instances of user.Username and returns an array of reviews of type *Review. 

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: 
<br />
type Review struct {
    gorm.Model
    GameName    string  //Names of game being reviewed
    Rating      float32 //Rating (out of 5) of the game
    Description string  //Description of the game played
    Username    string  //Name of the account
    PlayStatus  string  //PLAYING, DROPPED, COMPLETED, ON HOLD
}

<br />

8) GET: {Get a json of the 4 most recent games}: http://localhost:8080/recent

Functionality
<br />
This function returns a json of the 4 most recent games added to the games database. The function starts by using RAWG SDK’s NewGamesFilter() function along with SetPageSize(int numElements) with parameter 4 and SetOrdering(string order) with parameter released. The combination of these functions returns one page of games of size 4 in order of release starting with the most recent game added to the database. These games are then packaged using json.Marshal and sent to the header.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: N/A

Json Returned & Format: 
<br />
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

<br />

9) GET: {Get a json of the top 5 games based on QuestList user ratings}: 
http://localhost:8080/topgames

Functionality
<br />
The function of this method is to return the RAWG game information for the top 5 games that have the highest average user rating based on ratings from QuestList itself. Once a user makes a review, the average rating for a game based on all user input is recalculated and stored. This function opens the “UserGameRankings.db” database and utilizes Gorm’s “Where” function. It searches the database of all games’ average ratings and uses a quick sort algorithm to order the result in descending order based on ratings. The top 5 game names are searched using RAWG SDK’s NewGamesFilter(), SetPage() and SetSearch() functions using the game names and a page size of 1. The array of ordered RAWG games are then marshaled and packaged into a json file which is sent to the header.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: {An Array of:}
<br />
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

<br />

10) GET: {Get a json, of the requested size, of games that will release within the next month}: 
http://localhost:8080/upcominggames

Functionality
<br />
This function returns an array of RAWG games that have not been released yet, but will be available within the next month, starting tomorrow. This uses a simple time calculation to calculate the next month’s time frame and pass that information to the RAWG SDK. An SDK filter is used with the NewGameFilter(), SetPageSize(), SetDates(), and SetOrdering function to find games within the time frame that will release soon, ordered in terms of pre-existing anticipation of the game retrieved from RAWG itself. The array of games are then marshaled and packed into a json file and sent to the header.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: N/A

Json Returned & Format: {An Array of:}
<br />
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}

<br />


11) GET: {Returns a json file of an array of user struct’s whose names somewhat match the passed in name}: 
http://localhost:8080/getuser

Functionality
<br />
This function returns an array of User objects most similar to the username provided in the URL contents. The function begins by opening the user database, “currentUsers.db”, and pulling the username with parameter name “user” from the URL. GORM’s AutoMigrate is then called to format the database to accept the User{} struct. Next, GORM’s db.Where() function is called with functionality “LIKE” to find all users with a name similar to the username pulled from the URL previously and stores it in the array titled users of type []User also storing an error in the hasUsers variable of type err. Lastly, if there are no users with a name similar to the query parameter, hasUsers will store an error and http.StatusInternalServerError will be thrown. Otherwise, if there are users in the “users” array, those objects will be converted to a byte array using json.Marshal() and written to the header along with a http.StatusOK. Additionally, the function returns the users array to be used for Unit Testing.

Status Returned:
<br />
If Successful: http.StatusCreated
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: 
<br />
type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex"`
    Email    string `gorm:"uniqueIndex"`
    Password string
}

<br />


12) GET: {Returns a json file of an array of review structs added within the past month from most recent to least recent}: 
http://localhost:8080/recentreviews

Functionality
<br />
This function returns all of the reviews from the last month in order of recency. This function starts by opening the “Reviews.db” database and setting the time frame for the function starting from the current time (start) to a month ago(end). Next, GORM’s AutoMigrate is then called to format the database to accept the Review{} struct. GORM’s db.Where() function then uses the “>” functionality with the updated_at column to find all reviews with an updated time greater than a month ago. The reviews are then stored in latestReviews of type []Review and recentReviews of type err. If recentReviews has a non-nil error, a http.StatusInternalServerError is written to the header, else latestReviews is reversed using the reverseArray([]Review) function to order the reviews from most to least recent and latestReviews is converted into a json and written to the header along with the status http.StatusOK. Additionally, if latestReviews has elements in it, they are returned for Unit Testing, else nil is returned.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: 
<br />
type Review struct {
    gorm.Model
    GameName    string  //Names of game being reviewed
    Rating      float32 //Rating (out of 5) of the game
    Description string  //Description of the game played
    Username    string  //Name of the account
    PlayStatus  string  //PLAYING, DROPPED, COMPLETED, ON HOLD
}

<br />

13) GET: {Returns a json of the game with the most number of reviews in the last month}: 
http://localhost:8080/featuredgame 

Functionality
<br />
This function returns a RAWG game ame object that corresponds to the element in UserGameRankings.db with the most reviews. This function starts by opening the “UserGameRankings.db” and setting the time frame for the function starting from the current time (start) to a month ago(end). Next, GORM’s AutoMigrate is then called to format the database to accept the GameRanking{} struct. GORM’s db.Where() function then uses the “>” functionality with the updated_at column to find all reviews with an updated time greater than a month ago. The game rankings are then stored in gameRankings of type []GameRankings and recentRankings of type err. If recentRankings has a non-nil error, a http.StatusInternalServerError is written to the header. Otherwise, the name of the game with the highest number of reviews is stored in featuredGame of type string. The rawg function chain NewGamesFilter().SetPageSize(int pageSize).SetSearch(string gameName) is then used with parameters 1 and featuredGame respectively and stores a []*rawg.Game object in games. The array games is an array of size 1 containing the game with name featuredGame. For functionality purposes, this game is stored in variable “game” of type *rawg.Game by defining it with games[0]. This game is then converted to a json file and sent to the header along with the status http.StatusOK. Additionally, if game is defined, it is returned for Unit Testing, else nil is returned.

Status Returned:
<br />
If Successful: http.StatusOK
<br />
If Unsuccessful: http.StatusInternalServerError

Json Returned & Format: 
type Game struct {
	ID               int            `json:"id"`
	Slug             string         `json:"slug"`
	Name             string         `json:"name"`
	Released         DateTime       `json:"released"`
	Tba              bool           `json:"tba"`
	ImageBackground  string         `json:"background_image"`
	Rating           float32        `json:"rating"`
	RatingTop        int            `json:"rating_top"`
	Ratings          []*Rating      `json:"ratings"`
	RatingsCount     int            `json:"ratings_count"`
	ReviewsTextCount int            `json:"reviews_text_count"`
	Added            int            `json:"added"`
	AddedByStatus    *AddedByStatus `json:"added_by_status"`
	Metacritic       int            `json:"metacritic"`
	Playtime         int            `json:"playtime"`
	SuggestionsCount int            `json:"suggestions_count"`
	ReviewsCount     int            `json:"reviews_count"`
	SaturatedColor   string         `json:"saturated_color"`
	DominantColor    string         `json:"dominant_color"`
	Platforms        []*struct {
		Platform       *Platform    `json:"platform"`
		ReleasedAt     DateTime     `json:"released_at"`
		RequirementsEn *Requirement `json:"requirements_en"`
		RequirementsRu *Requirement `json:"requirements_ru"`
	} `json:"platforms"`
	ParentPlatforms []*struct {
		Platform struct {
			ID   int    `json:"id"`
			Slug string `json:"slug"`
			Name string `json:"name"`
		}
	} `json:"parent_platforms"`
	Genres []*Genre `json:"genres"`
	Stores []*struct {
		ID    int    `json:"id"`
		Store *Store `json:"store"`
		UrlEn string `json:"url_en"`
		UrlRu string `json:"url_ru"`
	} `json:"stores"`
	Clip             *Clip  `json:"clip"`
	Tags             []*Tag `json:"tags"`
	ShortScreenshots []*struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"short_screenshots"`
}


 
