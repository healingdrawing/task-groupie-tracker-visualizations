package server

import (
	"fmt"
	"groupie-tracker/fetch"
	"log"
	"net/http"
	"os"
)

var (
	ApiURL    string
	ImagesURL string
	ImagesDir string
	LogPath   string
)

const GreenOK = "\033[32mOK\033[0m"
const RedFAIL = "\033[31mFAIL\033[0m"

var accessLog *log.Logger

func init() {
	if ApiURL = os.Getenv("GROUPIE_API_URL"); ApiURL == "" {
		ApiURL = "https://groupietrackers.herokuapp.com/api"
	}
	if ImagesURL = os.Getenv("GROUPIE_IMAGES_URL"); ImagesURL == "" {
		ImagesURL = "/images/"
	}
	if ImagesDir = os.Getenv("GROUPIE_IMAGES_DIR"); ImagesDir == "" {
		ImagesDir = "./images"
	}
	if LogPath = os.Getenv("GROUPIE_LOG_PATH"); LogPath == "" {
		LogPath = "./access.log"
	}
}

/*
Start fetches data from API and returns http.Handler with API and images endpoints

You can customize API by changing following environment variables:

`GROUPIE_API_URL` - data source API, default is `https://groupietrackers.herokuapp.com/api`

`GROUPIE_IMAGES_DIR` - folder to locally store images from API, default is `./images`

`GROUPIE_IMAGES_URL` - URL to serve locally stored images, default is `/images/`
*/
func Start() http.Handler {
	lf, err := os.OpenFile(LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		log.Fatal("can not open access.log: ", err)
	}
	accessLog = log.New(lf, "", log.LstdFlags)
	fmt.Printf("Initilizing server with config: Images URL: '%v', Images Dir: '%v'\n", ImagesURL, ImagesDir)
	err = os.RemoveAll(ImagesDir)
	if err != nil {
		log.Fatal("can not clean folder for images: ", err)
	}

	err = os.MkdirAll(ImagesDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Fetching data from api...")
	artists, err := fetch.Artists(ApiURL)
	if err != nil {
		fmt.Println(RedFAIL)
		log.Fatal(err)
	}
	fmt.Println(GreenOK)

	fmt.Print("Downloading images...")
	err = fetch.Images(artists, ImagesDir, ImagesURL)
	if err != nil {
		fmt.Println(RedFAIL)
		log.Fatal(err)
	}
	fmt.Println(GreenOK)

	router := http.NewServeMux()

	// serve images
	router.HandleFunc(ImagesURL, imagesHandler)
	router.HandleFunc("/api/artists/", handle(artists).artistsHandler)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessLog.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		// TODO: add export mode with client-side only data fetching
		// (w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		router.ServeHTTP(w, r)
	})
}
