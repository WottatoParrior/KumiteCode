package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/sample", samplePage)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to KumiteCode \n")
	fmt.Fprintf(w, "The place to interview potential candidates using Code Review \n")
}

func samplePage(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("KumiteCode is up and running")
	DBConnect()
	handleRequests()
}
