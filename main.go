package main

import (
	"log"

	"github.com/joho/godotenv"
)

type Sources struct {
	Sources []Source `json: "sources"`
}
type Source struct {
	Id          string `json: "id"`
	Name        string `json: "name"`
	Description string `json: "description"`
	Category    string `json: "category"`
	Url         string `json: "url"`
	Language    string `json: "language"`
	Country     string `json: "country"`
}

type Articles struct {
	Articles     []Article `json: "articles"`
	TotalResults int       `json: "totalReults"`
}
type Article struct {
	Author      string `json: "author"`
	Title       string `json: "title"`
	Description string `json: "description"`
	Url         string `json: "url"`
	UrlToImage  string `json: "urlToImage"`
	PublishedAt string `json: "publishedAt"`
	Source      Source `json: "source"`
}

var PRODUCTION bool = true

func main() {

	err := godotenv.Load()
	if err != nil && PRODUCTION == false {
		log.Fatal("Error:", err)
	}
	createServer()

}
