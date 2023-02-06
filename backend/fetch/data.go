package fetch

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/format"
	"groupie-tracker/models"
	"io"
	"net/http"
	"sort"
	"strings"
)

// Artists fetches all data from API by url and saves it to array of models.Artist
func Artists(url string) ([]models.Artist, error) {
	var artists []models.Artist
	var urls struct {
		Artists   string `json:"artists"`
		Locations string `json:"locations"`
		Dates     string `json:"dates"`
		Relations string `json:"relation"`
	}
	err := fetchURL(url, &urls) // get the urls
	if err != nil {
		return nil, fmt.Errorf("can not fetch urls: %w", err)
	}

	err = fetchURL(urls.Artists, &artists) // getting all data from artists url
	if err != nil {
		return nil, fmt.Errorf("can not fetch artists: %w", err)
	}

	artistsMap := make(map[int]*models.Artist) // creating map to quickly get artist by id
	for i := range artists {
		artistsMap[artists[i].ID] = &artists[i]
	}

	var relations struct {
		Index []struct {
			ID             int                 `json:"id"`             // id of the artist
			DatesLocations map[string][]string `json:"datesLocations"` // map["the_city-the_country"][]string{"dd-mm-yyyy", ...}
		} `json:"index"`
	}
	err = fetchURL(urls.Relations, &relations)
	if err != nil {
		return nil, fmt.Errorf("can not fetch relations: %w", err)
	}
	for _, relation := range relations.Index {
		var concerts []models.City
		for location, dates := range relation.DatesLocations {
			for i := range dates {
				// format dates to ISO 8601 format (YYYY-MM-DD)
				dayMonthYear := strings.Split(dates[i], "-")
				if len(dates[i]) != len("DD-MM-YYYY") || len(dayMonthYear) != 3 {
					return nil, fmt.Errorf("wrong date format %v, expected DD-MM-YYYY", dates[i])
				}
				day, month, year := dayMonthYear[0], dayMonthYear[1], dayMonthYear[2]
				dates[i] = fmt.Sprintf("%v-%v-%v", year, month, day)
			}
			sort.Strings(dates)
			concerts = append(concerts,
				models.City{
					Name:  format.Location(location),
					Dates: dates,
				})
		}
		sort.Slice(concerts, func(i, j int) bool {
			return strings.Compare(concerts[i].Dates[0], concerts[j].Dates[0]) == -1
		})
		artistsMap[relation.ID].Cities = concerts
	}

	return artists, nil
}

// fetchURL fetches json by url and json.Unmarshal it to dst
func fetchURL[T any](url string, dst *T) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseData, dst)
	if err != nil {
		return err
	}
	return nil
}
