Javier Lopez, Richard Sorkin, Sean Cheema, and Thomas Gallego

2/23/23

Professor Dobra

**Sprint 2 Documentation**
<br/>

**_User Stories Worked On This Sprint:_**
<br />


**2) Game Library** <hr />As a **website visiter** I would like to **have access and see a list of all the games documented on the website** for the purpose of **knowing what I would like to add to my user library**.
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

**Work Completed**
<hr />

**Frontend:**

 
**Backend:**

Work Completed:
This spring focused on the communication between frontend and back end. We began by finding a way to link the two sides of the project. To achieve this, we used the URL handler and parameters passed into it to get directions from the front. Using gorilla mux, we defined a plethora of functions that each had a specified url that were each linked to a function that would be called when that handler was typed in. For example, a url with the handler “/allGames/1” would call a function that lists all games in page one of our database in json format. We began by defining handler functions for specific URL’s and defining  parameters in those handlers. Those parameters would be used to send the backend information based on the front end’s requirements. These parameters would be explicitly listed in our handler functions using curly braces “{}” and would be extracted in our functions using mux.Vars() methods.

The process begins by having the frontend invoke the URL’s we specified, putting the desired parameters based on our URL’s definitions. The backend listens to calls to certain URLs and once any handler is called that is defined in our main function, it calls its matching function which either communicates with our database and returns desired information (as a json) or modified and stored some information. To ensure that the front end can call and access the backend’s server, we had to add CORS in the header of the json files we return to the front end to prevent any security measures from limiting the two server’s communication.

In this sprint, we cleaned up a lot of the code. We added comments to each function declaration and each step within those functions to thoroughly follow the process. We also reformatted the way our existing functions worked. Our main functionality involve 3 methods that return json of certain types of games that the front end can interpret and display to the user. Our first function AllGames creates a filter, using our API, that returns a page’s worth of games. Each page contains only about 40 games. Each call to this function requires the handler to have a page parameter that the front end provides. Using this, we request that specified page from RAWG’s database and return the json with those games to the front. In the method Games, we do a similar process. We take in a parameter from the url for the name of the game that the user wants to search for. We extract this information, known as a slug, and search RAWG’s database for the 10 games that best match the specified name. We compact these games into a json file and send it to the front. Lastly, we created a function called RecentGames which similarly to the previous functions, returns a set of games. In this case, we do not take any parameters and simply search the database for the 4 most recently added games. We then compact and send the json to the frontend.

Back End Unit Tests:
We have implemented a variety of Unit Tests using the built-in Go testing. Initially, we created TestHello() to test the functionality of the Go “testing” import. To test our Main.go functions we created tests TestGame() and TestAllGames(). 

TestGame():
TestGame() tests the function Game() with a specific game “slug” (slug is a concatenated version of a game name setting all letters to lowercase and replacing all spaces with hyphens). When this slug is passed, Game should return the 10 most similar games to that slug with array[0] being the most similar and array[9] being the least similar. For this test we pulled the first element of the array as it is the desired game and compared it to our intended game string. 

TestAllGames():
Next, we used TestAllGames() to test the function AllGames(). AllGames() should return every game in the RAWG API one “page” at a time. Each page will have 40 games starting from the most popular game in the API to the least popular for all 800,000+ games. For this test, we pulled the first element of the first page which would be the most popular game in the API, in this case “Grand Theft Auto V”. 

Testing Limitations:
Lastly, it should be noted that our SignUp() and RecentGames() could not be tested. Our current SignUp() implementation is not currently set up to support users correctly and will be changed in a later sprint. As for RecentGames(), there is no way to predict the outcome of RecentGames() as new games are being added to the API every day causing the RecentGames() output to be changed frequently.

API Documentation:

The API we used is titled RAWG API. RAWG is a database of games storing a variety of information including but not limited to: name, image, different ratings, release date, and developers. We have used this data to organize and traverse the database to acquire desired information. In order to easily use the RAWG API, we found “RAWG SDK GO”, a client with built in functions to traverse the RAWG database. The documentation for both RAWG API and RAWG SDK GO are linked below. 

RAWG API: https://api.rawg.io/docs/
RAWG SDK GO: https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-documentation



