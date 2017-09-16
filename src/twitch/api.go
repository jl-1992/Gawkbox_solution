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

/*set up GetUser and Get Channel methods, look at Twitch API*/
//Is client_id same for all IP addresses??

func GetUserInfo(w http.ResponseWriter, r *http.Request, username string){
	type Fields struct {
		//Id   int64  `json:"_id"`
		Bio string `json:"bio"`
		CreatedAt string `json:"created_at"`
		DisplayName string `json:"display_name"`
	}
	res, err := http.Get("https://api.twitch.tv/kraken/users/"+username+"/?client_id=tsxcx1up34ibjx1ttvjjlpkq69t737")
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
	w.Header().Set("content-type", "application/json")
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
	w.Header().Set("content-type", "application/json")
	w.Write([]byte("Channel info:\n"))
	w.Write(output)
	w.Write([]byte("\n\n"))
}

func IsStreaming(w http.ResponseWriter, r *http.Request, username string){
	type Fields struct {
		Fields struct {
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

	var stream_map map[string]bool = make(map[string]bool)
	if fields.Fields.Video_Height==0{
		stream_map = map[string]bool{"stream": false}
	}else{
		stream_map = map[string]bool{"stream": true}
	}
	output, _ = json.MarshalIndent(stream_map, "", "  ")
	w.Write([]byte("Stream info:\n"))
	w.Header().Set("content-type", "application/json")
	w.Write(output)	
}