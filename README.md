Tech News Highlights
===================
## Description
A web app for showing user highlights from tech news websites - sources & their highlights, using [newsapi](https://newsapi.org).

I did this project to understand the process of building simple web apps and rest apis with Golang.

I believe its also possible build this web app with Angular only without the rest api but my end goal was learning Go.

-----------------------------------------------------------------------

## User Requirements

- A user is able to click on a tech news source and view a list of articles from that source.
- Clicking on the news article will then redirect you to the news article's web page.

## Quick Start

### Installation
- Install go
- `git clone https://github.com/jakhax/go-tech-news-highlights.git`
- `cd go-tech-news-highlights`
- Add your newsapi key to .env file `APIKEY=YOUR_KEY`, do commit this file.
- `go run main.go http_handlers.go` or   `go build`
- for angular `cd go-tech-news-highlights/front-end && npm i`

### Go Rest API endpoints

You can test the endpoints using a browser/curl/postman or any tool you prefer
### Tech news sources
`/sources`
### Source articles
`/articles/{article_is}`
### Source top headlines
`/top-headlines/{article_is}`

## Live Demo

- [Front end demo](https://jakhax.github.io/go-tech-news-highlights)
- [Heroku Go Restapi](https://tech-news-highlights.herokuapp.com/sources) - endpoints `/sources` or `/article/{article_id}` or `/top-headlines/{article_id}`

## Deployment
### Deploying back-end rest api to Heroku 

This assumes you have ever worked with heroku, not necessarily Go applications.
- Create a `Procfile` , a text file in the root directory of your application, to explicitly declare what command should be executed to start your app.
    - `web: go-tech-news-highlights`
- Install Go deps to manage Go package dependencies on Heroku, which helps build applications reproducibly by fixing their dependencies.
    - `go get github.com/tools/godep`
- Run `godep save ./...` to generate dependencies.
- Create Heroku app, commit all changes and deploy.
- Add your news `APIKEY` to heroku via the dashboard settings area or by running `heroku config:set APIKEY=YOUR_API_KEY`
- edit `PRODUCTION` variable to `true`.
- Test your app via browser / postman / curl , see `heroku logs` incase of an error.

### Deploying angular frontend to ghpages
- `npm i -g angular-cli-ghpages`
- `ng build --prod --base-href="https://GithubUserName.github.io/GithubREPO/"` to create a build folder `dist`, copy this folder to the root folder and run `ngh --dir=dist/`
- Remember to add ghpages url to back-end ResponseWriter http header to avoid issues with cors
    - `w.Header().Set("Access-Control-Allow-Origin", "https://jakhax.github.io")`

## Technologies Used
* [Golang](https://golang.org)
* [Angular 6](https://angular.io/)

## References
* [Golang Docs](https://golang.org)
* [News Api Docs](https://newsapi.org/docs)

## Bugs & Contribution
- I dont plan on supporting this or responding to issues concerning this project, this is just an assignment, while I learn basic Golang.

## License ([MIT License](http://choosealicense.com/licenses/mit/))
This project is licensed under the MIT Open Source license, (c) [Jack ogina](https://github.com/jakhax)
