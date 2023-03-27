Javier Lopez, Richard Sorkin, Sean Cheema, and Thomas Gallego

3/27/23

Professor Dobra

**Sprint 3 Documentation**
<br/>

**_User Stories Worked On This Sprint:_**
<br />


**1) Game Library** <hr />As a **website visiter** I would like to **have access and see a list of all the games documented on the website** for the purpose of **knowing what I would like to add to my user library**.
<br />
**Implementation Details** <hr /> 
- Finding a useable public database of most available games in the market to use as our back bone for the video games users can add to their libraries 
- Create a filter function that can search for games in the list in a time sensitive manner
- Implemented by front and back end 
- Front end handles game information display 
- Back end handles library importing and file traversal
<br />

**2) Home Page** <hr />As a **site visitor** I would like to **access a home page** for the purpose of **create a hub for navigation.**
<br />
**Implementation Details** <hr /> 
- Implemented by front and back-end
- Create a GUI to quickly access other pages
- Display new game information to site visitors
<br />

**3) Sign-Up** <hr />As a new user, I would like to sign up and create an account for the purpose of saving my activity on the site
<br />
**Implementation Details** <hr />
- Implemented by front and back end
- Create form gui
- Add users to user database in backend
<br />

**4) Login** <hr />As a returning user I would like to log into the website for the purpose of seeing my previous activity on the site
<br />
**Implementation Details** <hr />
- Implemented by front and back end
- Front end create login gui
- Back end get user information
<br />

**5) Review Issue** <hr />As a site visitor I would like to review games for the purpose of keeping track of the games I have played/am playing and how I liked them
<br />
**Implementation Details** <hr />
- Implemented by front and back end
- Create GUI to be able to review games
- Create a review database with elements GameName, Rating, Description, Username, and PlayStatus
- If a user creates a review for a game they have already reviewed, it will overwrite that review
<br />


**Work Completed**
<hr />

**Frontend:**
_Work Completed_:<br/> ~~~~~~~~~~~~~~~~~~~~~Detail Work Completed and How

_Units Tests_:<br/> ~~~~~~~~~~~~~~~~~~~~~~~~List and describe your unit tests and thier functions

_HttpClient:_<br/>
We started by focusing on implementing an HttpClient so that we could effectively communicate with the back-end. Once we we achieved this we were able to send JSON files between the front-end and the back-end.

_Game Service:_<br/>
We then created an injectable Game Service class with a function that gets an observable from the back-end with a parameter ***page***. This parameter specifies which page should be queried from the back-end. If the Game Service fails to obtain the observable from the back-end, an error handler function is called, and a back-end error code is returned.

_Game Component:_<br/>
We created a game component that implements the Game Service. The game component subscribes to the data returned from the aforementioned function and displays each game object using the NgFor tag in the component's HTML.

_Cypress Testing:_<br/>
Utilized Cypress to test that game component can be mounted and that it can get data from the back-end API.

_Summary:_ <br/>
We obtained an array of games from the back-end and displayed them on the top-games page. We were successful in connecting front and back-end.

-----------------------------------------------WORK IN PROGRESS-------------------------------------------

  <hr />

**Backend:**

_Work Completed_:<br/>
In this sprint, we worked on the User and Review databases and all of the functions associated with it. We started by making User and Review structs to store specific information about these objects. The user struct stores a unique username, unique email, and a password. For the review struct, it stores the game being reviewed, rating out of 5, description, username of the user who reviewed the game, and a playstatus .To retrieve these objects from the front-end, we pull body content from the front-end using json.Decode() and format these objects into the struct. The main functionality that we changed in this sprint regarding back-end urls is how we read in parameters. Originally, we were pulling parameters via mux.Vars() which pulls the parameters straight from the url (e.g. /specific-game/overwatch), but now we are using the url query parameters to specify needed attributes (e.g. /games?page=1). To send, we used the json.Marshal() function and wrote to the header. As for the functionality associated with these databases, we are now able to create user and review structs, add them to the database, and look them up.
<br/>

_Back End Unit Tests_: <br/>
We have implemented a variety of Unit Tests using the built-in Go testing. Initially, we created TestHello() to test the functionality of the Go “testing” import. To test our Main.go functions we created tests TestGame(), TestAllGames(), TestSignUp(), TestSignIn(), TestWriteReview(), and TestGetReview(). 

-TestGame():<br/>
TestGame() tests the function Game() with a specific game “slug” (slug is a concatenated version of a game name setting all letters to lowercase and replacing all spaces with hyphens). When this slug is passed, Game should return the 10 most similar games to that slug with array[0] being the most similar and array[9] being the least similar. For this test we pulled the first element of the array as it is the desired game and compared it to our intended game string. 

-TestAllGames():<br/>
Next, we used TestAllGames() to test the function AllGames(). AllGames() should return every game in the RAWG API one “page” at a time. Each page will have 40 games starting from the most popular game in the API to the least popular for all 800,000+ games. For this test, we pulled the first element of the first page which would be the most popular game in the API, in this case “Grand Theft Auto V”. 

-TestSignUp():<br/>
This tests the functionality of Sign-Up(). This tested function should add a new user to the database of users if their username and/or email have not been used already. For this test, we created an arbitary username, password, and email and inserted them into the database. To ensure they were added, we scan the database to find the first, and only occurrance of their username. To test the prevention of duplicate users, we try to add an already existing username and expect an error.

-TestSignIN():<br/>
This tests the functionality of Sign-In(). Sign-In() should check that a requested username exists in the database. If they do exisit, we compare the stored password with the requested one. If everything matches up, we return an OK status and mark the signed-in user as the active user. If something doesn't match the database, we return an error status. Our test first tries to find a user that does exist with a mathcing password. This should return an OK status. Afterwards, it tries logging in with a nonexistent user, resulting in an error.

-TestWriteReview():<br/>
This tests the function WriteAReview() with a review struct passed in containing elements such as the game being reviewed, rating out of 5, description, username of the user who reviewed the game, and a playstatus. The intended functionality is that WriteAReview() will add the review to the Review database if that user has not already reviewed that game, and modify the review if a user has already created a review for that game.

-TestGetReview():<br/>
This tests the function GetReview() with a user struct passed in containing elements such as a unique username, unique email, and a password. The intended functionality of GetReview() is that the function retrieves all instances of a specific user’s reviews from the review database and writes an error status to the header if the user has not created a review.

-Testing Limitations:<br/>
It should be noted that RecentGames() could not be tested. For RecentGames(), there is no way to predict the outcome of RecentGames() as new games are being added to the API every day causing the RecentGames() output to be changed frequently.

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






 
