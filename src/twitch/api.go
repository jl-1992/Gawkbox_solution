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
	type Fields struct {
		//Id   int64  `json:"_id"`
		Bio string `json:"bio"`
		CreatedAt string `json:"created_at"`
		DisplayName string `json:"display_name"`
		//Name string `json:"name"`
	}
	res, err := http.Get("https://api.twitch.tv/kraken/users/"+username+"/?client_id=tsxcx1up34ibjx1ttvjjlpkq69t737")
	if err != nil{
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	//defer r.Body.Close()
	var fields Fields
	err = json.Unmarshal(body, &fields)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(fields)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}