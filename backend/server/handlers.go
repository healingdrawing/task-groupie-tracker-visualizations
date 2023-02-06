package server

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type handle []models.Artist

func (artists handle) artistsHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			errorResponse(w, http.StatusInternalServerError)
		}
	}()

	if r.Method != http.MethodGet {
		errorResponse(w, http.StatusMethodNotAllowed)
		return
	}

	sendObject := func(object any) {
		w.Header().Set("Content-Type", "application/json")
		artistsJson, err := json.Marshal(object)
		if err != nil {
			panic(err) // this will call recover with error 500
		}
		_, err = w.Write(artistsJson)
		if err != nil {
			log.Print(err)
		}
	}

	if r.URL.Path == "/api/artists/" {
		params := r.URL.Query()
		if len(params) == 0 {
			// TODO: for bigger amount of artists it's better to add padding and send only partial info (id, name and image)
			sendObject(artists)
			return
		}
		for _, artist := range artists {
			if artist.Name == params.Get("name") ||
				strconv.Itoa(artist.ID) == params.Get("id") {
				sendObject(artist)
				return
			}
		}
		errorResponse(w, http.StatusNotFound)
		return
	}

	idStr := strings.TrimPrefix(strings.TrimSuffix(r.URL.Path, "/"), "/api/artists/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorResponse(w, http.StatusNotFound)
		return
	}

	for _, artist := range artists {
		if artist.ID == id {
			sendObject(artist)
			return
		}
	}
	// there's no such artist if we're here
	errorResponse(w, http.StatusNotFound)
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	imagePath := strings.TrimPrefix(r.URL.Path, ImagesURL)
	if imagePath == "" {
		errorResponse(w, http.StatusNotFound)
		return
	}
	http.ServeFile(w, r, path.Join(ImagesDir, imagePath))
}

func errorResponse(w http.ResponseWriter, code int) {
	// default error text is just "Not Found"
	http.Error(w, fmt.Sprintf("%v %v", code, http.StatusText(code)), code)
}
