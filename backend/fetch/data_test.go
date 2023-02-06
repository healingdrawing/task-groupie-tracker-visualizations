package fetch

import (
	"groupie-tracker/models"
	"log"
	"net/url"
	"testing"
)

func notEmptyName(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		if artist.Name == "" {
			t.Errorf("empty name of the artist with ID [%d]\n", artist.ID)
		}
	})
}

func notEmptyID(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		if artist.ID < 1 {
			t.Errorf("not positive ID of the artist [%s]\n", artist.Name)
		}
	})
}

func notEmptyCitiesAndDates(artist models.Artist, t *testing.T) {
	for _, city := range artist.Cities {
		t.Run("", func(t *testing.T) {
			if city.Name == "" {
				t.Errorf("empty city name\n")
			}
		})
		for _, day := range city.Dates {
			t.Run("", func(t *testing.T) {
				if day == "" {
					t.Errorf("empty date\n")
				}
			})
		}
	}
}

func notEmptyCreationDate(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		if artist.CreationDate < 1 {
			t.Errorf("not positive creation date [%d] of the artist [%s]\n", artist.CreationDate, artist.Name)
		}
	})
}

func notEmptyFirstAlbum(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		if artist.FirstAlbum == "" {
			t.Errorf("empty first album name of the artist [%s]\n", artist.Name)
		}
	})
}

func notEmptyImageURL(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		if artist.ImageURL == "" {
			t.Errorf("empty image url of the artist [%s]\n", artist.Name)
		}
	})
}

func notEmptyMembers(artist models.Artist, t *testing.T) {
	t.Run("", func(t *testing.T) {
		for _, member := range artist.Members {
			if member == "" {
				t.Errorf("empty member of the artist [%s]\n", artist.Name)
			}
		}
	})
}

func TestData(t *testing.T) {
	log.Printf("\nInternet connection required for testing!!!\n\n")

	rawData, err := Artists("https://groupietrackers.herokuapp.com/api")
	if err != nil {
		t.Error(err)
	}

	for _, artist := range rawData {
		notEmptyName(artist, t)
		notEmptyID(artist, t)
		notEmptyCitiesAndDates(artist, t)
		notEmptyCreationDate(artist, t)
		notEmptyFirstAlbum(artist, t)
		notEmptyImageURL(artist, t)
		notEmptyMembers(artist, t)
	}

	t.Run("", func(t *testing.T) {
		var dummy struct {
			Artists   string `json:"artists"`
			Locations string `json:"locations"`
			Dates     string `json:"dates"`
			Relation  string `json:"relation"`
		}
		err := fetchURL("https://groupietrackers.herokuapp.com/api", &dummy)
		if err != nil {
			t.Error(err)
		}

		{
			_, err := url.ParseRequestURI(dummy.Artists)
			if err != nil {
				t.Error(err)
			}
		}

		{
			_, err := url.ParseRequestURI(dummy.Dates)
			if err != nil {
				t.Error(err)
			}
		}

		{
			_, err := url.ParseRequestURI(dummy.Locations)
			if err != nil {
				t.Error(err)
			}
		}

		{
			_, err := url.ParseRequestURI(dummy.Relation)
			if err != nil {
				t.Error(err)
			}
		}
	})

	// comparison testing of fetched data with alternative sources with similar data

	t.Run("", func(t *testing.T) {
		{
			var dummyDates struct {
				Index []struct {
					ID    int      `json:"id"`
					Dates []string `json:"dates"`
				} `json:"index"`
			}
			err := fetchURL("https://groupietrackers.herokuapp.com/api/dates", &dummyDates)
			if err != nil {
				t.Error(err)
			}
			if len(dummyDates.Index) != len(rawData) {
				t.Error("The length of the fetched data is not equal. Potential lost of the data!")
			}
		}

		{
			var dummyLocations struct {
				Index []struct {
					ID        int      `json:"id"`
					Locations []string `json:"locations"`
					Dates     string   `json:"dates"`
				} `json:"index"`
			}

			err := fetchURL("https://groupietrackers.herokuapp.com/api/locations", &dummyLocations)
			if err != nil {
				t.Error(err)
			}
			if len(dummyLocations.Index) != len(rawData) {
				t.Error("The length of the fetched data is not equal. Potential lost of the data!")
			}
		}

	})
}
