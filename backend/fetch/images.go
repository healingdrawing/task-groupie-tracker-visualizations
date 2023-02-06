package fetch

import (
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// Images fetches all images by URLs provided in models.Artist
// it stores them locally in saveToDir and changes path in models.Artist to relative imagesURL
func Images(artists []models.Artist, saveToDir string, imagesURL string) error {
	for i := range artists {
		_, imageFileName := filepath.Split(artists[i].ImageURL)
		err := fetchImage(artists[i].ImageURL, filepath.Join(saveToDir, imageFileName))
		if err != nil {
			return err
		}
		imageURL, err := url.JoinPath(imagesURL, imageFileName)
		if err != nil {
			return err
		}
		artists[i].ImageURL = imageURL
	}
	return nil
}

// fetchImage fetches image by url and saves it to the specified path
func fetchImage(URL, saveToPath string) error {
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK &&
		!strings.HasPrefix(response.Header.Get("Content-Type"), "image/") {
		return fmt.Errorf("reponse is not a successful image response:\n%+v", response)
	}
	//Create an empty file
	file, err := os.Create(saveToPath)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
