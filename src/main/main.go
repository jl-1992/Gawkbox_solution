package main

import (
	"fmt"
	"twitch"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Booting the server...")

	http.HandleFunc("/users/", myHandlerFunc)

	http.ListenAndServe(":8080", nil)
}

// myHandlerFunc - A handler function for the route /users for the HTTP server
func myHandlerFunc(w http.ResponseWriter, r *http.Request) {
	//If request method is anything but GET, throw error
	if r.Method != "GET" {
		http.Error(w, "Invalid request method.", 405)
	}else{
		//Trim URL path to get suffix of users/USERNAME so user=USERNAME
		user := strings.TrimPrefix(r.URL.Path, "/users/")
		w.Header().Set("content-type", "application/json")
		twitch.GetUserInfo(w,r,user)
		twitch.GetChannelInfo(w,r,user)
		twitch.IsStreaming(w,r,user)
	}
}
