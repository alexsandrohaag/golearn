package main

import (
	"fmt"
	"log"
	"net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Articles []Article


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}


func testPostArticles(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Test Post Article")
	fmt.Println("Endpoint Hit: testPostArticles")
}  


func returnAllArticles(w http.ResponseWriter, r *http.Request){
    articles := Articles{
        Article{Title:"Hello", Desc:"Article Description", Content:"Article Content"},
        Article{Title:"Hello 2", Desc:"Article Description 2", Content:"Article Content 2"},
    }
    fmt.Println("Endpoint Hit: returnAllArticles")
    
    json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]
    
    fmt.Fprintf(w, "Key: " + key)
}


func handleRequests() {
    
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/all", returnAllArticles).Methods("GET")
    myRouter.HandleFunc("/all",testPostArticles).Methods("POST")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":3000",myRouter))
}


func main() {
	handleRequests()
}
