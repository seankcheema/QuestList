Javier Lopez, Richard Sorkin, Sean Cheema, and Thomas Gallego

2/3/23

Professor Dobra

**Sprint 1 Documentation**
<br/>

**_User Stories:_**
<br />

**1) About Us** <hr />
As a **new site visitor** I would like to **access an about us page** for the purpose of **learning more about a website.**
<br />
**Implementation Details** <hr /> 
- Add a description of the website in the about us html file
- Style the page
- Implemented by front-end
<br />

**2) Sign Up** <hr />As a **new user**, I would like to **sign up and create an account** for the purpose of **saving my activity on the site**
<br />
**Implementation Details** <hr /> 
- Implemented by front and back end
- Create form gui
- Add users to user database in backend
<br />

**3) Log In** <hr />As a **returning user** I would like to **log into the website** for the purpose of **seeing my previous activity on the site**
<br />
**Implementation Details** <hr /> 
- Implemented by front and back end
- Front end create login gui
- Back end get user information
<br />

**4) Game Library** <hr />As a **website visiter** I would like to **have access and see a list of all the games documented on the website** for the purpose of **knowing what I would like to add to my user library**.
<br />
**Implementation Details** <hr /> 
- Finding a useable public database of most available games in the market to use as our back bone for the video games users can add to their libraries 
- Create a filter function that can search for games in the list in a time sensitive manner
- Implemented by front and back end 
- Front end handles game information display 
- Back end handles library importing and file traversal
<br />

**5) Home Page** <hr />As a **site visitor** I would like to **access a home page** for the purpose of **create a hub for navigation.**
<br />
**Implementation Details** <hr /> 
- Implemented by front and back-end
- Create a GUI to quickly access other pages
- Display new game information to site visitors
<br />

**_Issues to Address_**
<hr />

**Frontend:**

 - Communicate with backend server 
 - Display information from backend server on site 
 - Create a sign up/log in GUI Creating routing to navigate pages 
 - Implement Angular Materials to create GUI

Most of the time spent by the front-end team was learning how to get started with Angular and using the HTTPClientModule to communicate with the backend server. The frontend team was unfamiliar with the Angular framework and typescript prior to this project. The frontend team learned how to use Angular Materials in order to create a working GUI on the site and routing to move between pages.

**Backend:**

 - List item
 - Find working library of games 
 - Find database of games and their global rankings 
 - Create Game struct to store game data in our own library 
 - Be able to regurgitate game data in a readable way 
 - Find out how to link back end and front end 
 - Create user struct Each user has unique usernames and their own password 
 - Create database of users
 - Find a way to avoid creating a new client during each call to the handle functions

A large portion of the backend team’s time was spent learning about GoLang, Gorilla Mux, and Gorm. The language was new to both developers so extensive research was done learning the lexical structure of the language and the ways of writing, importing, and implementing code. Once acquainted with the language, our main concern shifted to getting a basic framework for our project. Since most of our website relies on the database and the items that that database brings (such as information on game names, publishers, launch dates, durations, pictures, etc) a lot of our efforts were spent finding a suitable database. After extensive searching, we found RAWG.io which required an access key which we were unable to access for some time. As a result, we extracted a database created on GitHub (found here: [https://github.com/bbrumm/databasestar/tree/main/sample_databases/sample_db_videogames](https://github.com/bbrumm/databasestar/tree/main/sample_databases/sample_db_videogames)) to playtest database access. We couldn’t find a suitable nor effective way to access the GitHub database and realized that a locally hosted database would be both ineffective and storage intensive.

  

We messaged the creators of the RAWG database and acquired an access key for use of the database. This allowed us to access a comprehensive library of games with countless elements in json files with data. To help link the database to our code, we found a library to help us interact with RAWG API (found here: [https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-readme](https://pkg.go.dev/github.com/dimuska139/rawg-sdk-go#section-readme)) that simplified the connection and added function that made creating a client and accessing elements from the database a lot simpler. We created a basic skeleton code that sets up a connection, requests and receives a complete list of over 800,000 games, and prints them out. This was done to assure access to the database and its content.

  

In short, we were able to find a usable database that included all features we needed. We were able to successfully link it to our code and we were also able to set up a server using gorilla mux that can listen for http requests from the front end. However, we have still not created the user struct to store individual user’s completed games, usernames, and passwords. This was due to the time we took setting up the database. It took longer than expected, but plan to have this done and created very early on in Sprint 2.
