package models

type Artist struct {
	ID           int      `json:"id"`
	ImageURL     string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Cities       []City   `json:"cities"`
}

type City struct {
	Name  string   `json:"name"`
	Dates []string `json:"dates"`
}
