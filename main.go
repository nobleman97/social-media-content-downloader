package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type showMe struct {
	Vidlink []string `json:"video"`
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getTicTok(w http.ResponseWriter, r *http.Request) {
	urlFromUser := r.FormValue("inputValue")
	addon := "https://tiktok-downloader-download-tiktok-videos-without-watermark.p.rapidapi.com/vid/index?url="

	url := fmt.Sprint(addon + urlFromUser)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "61e2ec60a6msh4f819358dd504ccp1d4de3jsnc408dd89b4bd")
	req.Header.Add("x-rapidapi-host", "tiktok-downloader-download-tiktok-videos-without-watermark.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	checkErr(err)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var holdData showMe

	err = json.Unmarshal(body, &holdData)
	checkErr(err)

	html, err := template.ParseFiles("resultpage.html")
	checkErr(err)

	giver := holdData

	html.ExecuteTemplate(w, "TikTok", giver)

}

func loadHome(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("index.html")
	checkErr(err)

	html.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", loadHome)
	http.HandleFunc("/vid_download/", getTicTok)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
