package main

import (
	"fmt"
	"twitch"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Booting the server...")

	// Configure a sample route
	http.HandleFunc("/users/", myHandlerFunc)

	// Run your server
	http.ListenAndServe(":8080", nil)
}

// myHandlerFunc - A sample handler function for the route /sample_route for your HTTP server
func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method.", 405)
	}
	user := strings.TrimPrefix(r.URL.Path, "/users/")
	twitch.GetUserInfo(w,r,user)
}
