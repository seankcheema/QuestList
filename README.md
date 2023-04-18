Welcome to QuestList!
--
QuestList is an open-source social media platform with the goal of creating a place where people of similar interests can come together to keep track of, rank, and share games they have played and enjoyed.

Getting Started
--
**Setting up the front-end:**

 1. Clone the repository in your IDE of choice. Visual studio code is recommended, but not necessary.
 2. Install *Node.js*. This is needed for the npm package manager and can be found at https://www.nodejs.org/en/download.
 3. Install the Angular CLI using the command: **npm install -g @angular/cli**. This is necessary to run the front-end and can also be used to easily create components, services, and other files.
 4. Finally, run the front-end server on your localhost by typing the command **ng serve** in your project directory. To view the front-end simply use navigate to **localhost:4200** in your web browser of choice.
 5. Additionally, the front-end unit test suite can be run by simply typing the **ng test** command.
 6.  The Cypress test suite can be installed using the npm package installer. Type the command **npm  install cypress --save-dev** to install Cypress. Then, to run the Cypress test suite simply type **npx cypress open**.

**Setting up the back-end:**

 1. Install **minGW-w64** (64-bit compiler). This can first be accomplished by installing, **MSYS2**, a collection of tools and libraries that provides up to date native builds of minGW-w64 and other software. More details on installing minGW-w64 can be found at: https://code.visualstudio.com/docs/cpp/config-mingw.
 2. Install **Go**. This is necessary for running back-end files and details on installation can be found at: https://go.dev/doc/install.
 3. Finally, run the back-end server on your localhost by first running *cd src/back-end* in your project directory. Then simply type **go run .** (the period is included). This will run all of the back-end go files!
 4. To manually view back-end handles simply navigate to **localhost:8080** in your browser and add the handle of choice.
 5. The back-end test suite can be run by simply running tests in the **Main_test.go** file!

Contributing
--
We encourage you to contribute to our project! To contribute simply clone the repository, code, and create a pull request. Note that in order for code to be merged into the main branch it needs to pass the current test suite and also be covered by new tests.

