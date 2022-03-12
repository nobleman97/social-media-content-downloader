package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	errEnv := godotenv.Load(".env")
	checkErr(errEnv)
	var apiKeyHolder = os.Getenv("API_KEY")

	urlFromUser := r.FormValue("inputValue")
	addon := "https://tiktok-downloader-download-tiktok-videos-without-watermark.p.rapidapi.com/vid/index?url="

	url := fmt.Sprint(addon + urlFromUser)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", apiKeyHolder)
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

	/* When DEPLOYING, I will use...
	
	port := os.Getenv("PORT")*/

	port := "8080"

	http.HandleFunc("/", loadHome)
	http.HandleFunc("/vid_download/", getTicTok)

	err := http.ListenAndServe(":"+port, nil)
	checkErr(err)

}
