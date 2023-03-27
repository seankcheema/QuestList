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


_Back End Unit Tests_: <br/>
We have implemented a variety of Unit Tests using the built-in Go testing. Initially, we created TestHello() to test the functionality of the Go “testing” import. To test our Main.go functions we created tests TestGame(), TestAllGames(), TestSignUp, TestSignIn, TestWriteReview, and TestGetReview. 

-TestGame():<br/>
TestGame() tests the function Game() with a specific game “slug” (slug is a concatenated version of a game name setting all letters to lowercase and replacing all spaces with hyphens). When this slug is passed, Game should return the 10 most similar games to that slug with array[0] being the most similar and array[9] being the least similar. For this test we pulled the first element of the array as it is the desired game and compared it to our intended game string. 

-TestAllGames():<br/>
Next, we used TestAllGames() to test the function AllGames(). AllGames() should return every game in the RAWG API one “page” at a time. Each page will have 40 games starting from the most popular game in the API to the least popular for all 800,000+ games. For this test, we pulled the first element of the first page which would be the most popular game in the API, in this case “Grand Theft Auto V”. 

-TestSignUp
TestSignUp():<br/>
This tests the functionality of Sign-Up(). This tested function should add a new user to the database of users if their username and/or passwords 

-Testing Limitations:<br/>
It should be noted that RecentGames() could not be tested. For RecentGames(), there is no way to predict the outcome of RecentGames() as new games are being added to the API every day causing the RecentGames() output to be changed frequently.

-External Library Documentation_:<br/>
The API we used is titled RAWG API. RAWG is a database of games storing a variety of information including but not limited to: name, image, different ratings, release date, and developers. We have used this data to organize and traverse the database to acquire desired information. In order to easily use the RAWG API, we found “RAWG SDK GO”, a client with built in functions to traverse the RAWG database. The documentation for both RAWG API and RAWG SDK GO are linked below. 

RAWG API: https://api.rawg.io/docs/
RAWG SDK GO: https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-documentation



