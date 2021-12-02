package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage! My name is Delaram")
	fmt.Println("Endpoint Hit: homePage")
}

func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TestPost endpoint worked")
}

// lest us do things like specyfy verbs we can use to access what endpoints in a far simpler fashion

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// we want it to be only accessible using get request...
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	// The call to json.NewEncoder(w).Encode(article) does the job of encoding our articles array into a JSON string and then writing as part of our response.
	json.NewEncoder(w).Encode(Articles)
}

func main() {

	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}
