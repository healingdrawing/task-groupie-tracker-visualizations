package server

import (
	"bytes"
	"encoding/json"
	"groupie-tracker/models"
	"groupie-tracker/server"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	router := server.Start()
	srv := httptest.NewServer(router)
	cli := srv.Client()

	var artists []models.Artist

	t.Run("[GET]/api/artists/", func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusOK)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		err = json.Unmarshal(body, &artists)
		if err != nil {
			t.Fatal("response is not a valid array of artists", err)
		}
		if len(artists) < 5 {
			t.Error("not enough artists in response: ", len(artists))
		}
		for _, artist := range artists {
			t.Run("[GET]"+artist.ImageURL, func(t *testing.T) {
				resp, err := cli.Get(srv.URL + artist.ImageURL)
				if err != nil {
					t.Fatal(err)
				}
				if resp.StatusCode != http.StatusOK {
					t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusOK)
				}
				contentType := resp.Header.Get("Content-Type")
				if !strings.HasPrefix(contentType, "image") {
					t.Fatalf("got content type %v, expected image...", contentType)
				}
			})
		}
	})

	t.Run("[GET]/api/artists/"+strconv.Itoa(artists[0].ID), func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/" + strconv.Itoa(artists[0].ID))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusOK)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		reference, err := json.Marshal(artists[0])
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(body, reference) != 0 {
			t.Errorf("got %s, expected %s", body, reference)
		}
	})
	t.Run("[GET]/api/artists/?name="+artists[0].Name, func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/?name=" + artists[0].Name)
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusOK)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		reference, err := json.Marshal(artists[0])
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(body, reference) != 0 {
			t.Errorf("got %s, expected %s", body, reference)
		}
	})
	t.Run("[GET]/api/artists/?id="+strconv.Itoa(artists[0].ID), func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/?id=" + strconv.Itoa(artists[0].ID))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusOK)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		reference, err := json.Marshal(artists[0])
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(body, reference) != 0 {
			t.Errorf("got %s, expected %s", body, reference)
		}
	})
	t.Run("[GET]/api/artists/CAT", func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/CAT")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusNotFound)
		}
	})
	t.Run("[GET]/api/artists/-1", func(t *testing.T) {
		resp, err := cli.Get(srv.URL + "/api/artists/-1")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != http.StatusNotFound {
			t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusNotFound)
		}
	})
	tests404 := []string{
		"/HELLO",
		"/",
		"/api/artists/-1",
		"/api/artists/CAT",
		"/api/artists/?dog=cat",
		"/api/artists/?name=SUPER_DUMMY_CAT_*(&*(*&*(",
		"/api/artists/?id=-1",
	}
	for _, test := range tests404 {
		t.Run("[GET]"+test, func(t *testing.T) {
			resp, err := cli.Get(srv.URL + test)
			if err != nil {
				t.Fatal(err)
			}
			if resp.StatusCode != http.StatusNotFound {
				t.Fatalf("got code %v, expected code %v", resp.StatusCode, http.StatusNotFound)
			}
		})
	}
}
