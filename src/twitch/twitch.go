package twitch

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func init() {
	fmt.Println("Initializing Twitch API...")
}

func GetUserInfo(w http.ResponseWriter, r *http.Request, username string){
	//Created a struct designating the fields we want that will be available from API
	type Fields struct {
		Bio string `json:"bio"`
		CreatedAt string `json:"created_at"`
		DisplayName string `json:"display_name"`
	}

	//received client_id when registering app and can use that to access Twitch API
	res, err := http.Get("https://api.twitch.tv/kraken/users/"+username+"/?client_id=tsxcx1up34ibjx1ttvjjlpkq69t737")
	if err != nil{
		panic(err.Error())
	}
	//Read all of the body from the response of the get request
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	var fields Fields
	//Unmarshal JSON from body and put into fields
	err = json.Unmarshal(body, &fields)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//Marshal JSON to /users/ HTTP page
	output, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("User info:\n"))
	w.Write(output)
	w.Write([]byte("\n\n"))
}

func GetChannelInfo(w http.ResponseWriter, r *http.Request, username string){
	type Fields struct {
		Views   int64  `json:"views"`
		Followers   int64  `json:"followers"`
		Game string `json:"game"`
		Language string `json:"language"`
	}
	res, err := http.Get("https://api.twitch.tv/kraken/channels/"+username+"/?client_id=tsxcx1up34ibjx1ttvjjlpkq69t737")
	if err != nil{
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	var fields Fields
	err = json.Unmarshal(body, &fields)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write([]byte("Channel info:\n"))
	w.Write(output)
	w.Write([]byte("\n\n"))
}

func IsStreaming(w http.ResponseWriter, r *http.Request, username string){
	type Fields struct {
		Fields struct {
			/*I used video_height as a field because a stream that is live
			returns an object "stream" which itself is an object. Thus,
			if the user is not streaming, video_height will be zero, but if 
			they are streaming, it will not be zero*/
			Video_Height int64 `json:"video_height"`
		} `json:"stream"`
	}
	res, err := http.Get("https://api.twitch.tv/kraken/streams/"+username+"/?client_id=tsxcx1up34ibjx1ttvjjlpkq69t737")
	if err != nil{
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()
	var fields Fields
	err = json.Unmarshal(body, &fields)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//Created my own JSON object to make stream key, value more basic
	var stream_map map[string]bool = make(map[string]bool)
	if fields.Fields.Video_Height==0{
		stream_map = map[string]bool{"stream": false}
	}else{
		stream_map = map[string]bool{"stream": true}
	}
	output, _ = json.MarshalIndent(stream_map, "", "  ")
	w.Write([]byte("Stream info:\n"))
	w.Write(output)	
}