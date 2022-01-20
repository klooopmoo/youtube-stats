package main

import (
	"fmt"
	"log"
	"net/http"
)

//Homepage- will be a simple hello world style page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World- Gabe Edition")

}

//setuproutes handles setting up our server
// routes and matching them to their respective functions
func setupRoutes() {
	http.HandleFunc("/", homePage)
	// we kick off our server to local host 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func main() {
	fmt.Println("YouTube Subscriber Monitor")
	setupRoutes()
}
