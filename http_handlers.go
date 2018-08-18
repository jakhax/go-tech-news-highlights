package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleGetSources(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if PRODUCTION {
		w.Header().Set("Access-Control-Allow-Origin", "https://jakhax.github.io/go-tech-news-highlights")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4200")
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	apiKey, ok := os.LookupEnv("APIKEY")
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Fatal("Unable to retrieve Api Key from environment vars")
	}
	sourcesLink := "https://newsapi.org/v2/sources?language=en&category=technology&apiKey=" + apiKey
	resp, err := http.Get(sourcesLink)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}

	var srcs Sources

	if err := json.Unmarshal(b, &srcs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}

	fmt.Println(srcs)
	json.NewEncoder(w).Encode(srcs)

	return

}
func handleGetSourceArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if PRODUCTION {
		w.Header().Set("Access-Control-Allow-Origin", "https://jakhax.github.io/go-tech-news-highlights")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4200")
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	apiKey, ok := os.LookupEnv("APIKEY")
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Fatal("Unable to retrieve Api Key from environment vars")
	}
	urlParams := r.URL.Query()
	page, ok := urlParams["page"]
	if !ok {
		page = []string{"1"}
	}
	params := mux.Vars(r)
	source_id := params["source_id"]
	articlesLink := "https://newsapi.org/v2/everything?language=en&sources=" + source_id + "&apiKey=" + apiKey + "&page=" + page[0]
	resp, err := http.Get(articlesLink)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}
	fmt.Println(string(b))
	var articles Articles

	if err := json.Unmarshal(b, &articles); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}

	fmt.Println(articles)
	json.NewEncoder(w).Encode(articles)

	return

}

func handleGetSourceTopHeadlines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if PRODUCTION {
		w.Header().Set("Access-Control-Allow-Origin", "https://jakhax.github.io/go-tech-news-highlights")
	} else {
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:4200")
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET")
	apiKey, ok := os.LookupEnv("APIKEY")
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Fatal("Unable to retrieve Api Key from environment vars")
	}
	urlParams := r.URL.Query()
	page, ok := urlParams["page"]
	if !ok {
		page = []string{"1"}
	}
	params := mux.Vars(r)
	source_id := params["source_id"]
	articlesLink := "https://newsapi.org/v2/top-headlines?language=en&sources=" + source_id + "&apiKey=" + apiKey + "&page=" + page[0]
	resp, err := http.Get(articlesLink)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}
	fmt.Println(string(b))
	var articles Articles

	if err := json.Unmarshal(b, &articles); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		log.Println("Error:", err)
		return
	}

	fmt.Println(articles)
	json.NewEncoder(w).Encode(articles)

	return

}

func createServer() {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/sources", handleGetSources).Methods("GET")
	muxRouter.HandleFunc("/articles/{source_id}", handleGetSourceArticles).Methods("GET")
	muxRouter.HandleFunc("/top-headlines/{source_id}", handleGetSourceTopHeadlines).Methods("GET")
	port := "8000"
	if PRODUCTION {
		port = os.Getenv("PORT")
	}
	s := http.Server{
		Addr:    ":" + port,
		Handler: handlers.LoggingHandler(os.Stdout, muxRouter),
	}
	log.Fatal(s.ListenAndServe())
}
