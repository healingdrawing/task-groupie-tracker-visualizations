# GROUPIE-TRACKER-BACKEND

RESTful API server to get data about artists and their concerts

---

Authors: [@maximihajlov](https://github.com/maximihajlov), [@healingdrawing](https://github.com/healingdrawing)
, [@nattikim](https://github.com/nattikim)

Solved during studying in Gritlab coding school on Åland, November 2022

---

## [Task description and audit questions](https://github.com/01-edu/public/tree/master/subjects/groupie-tracker)

## Usage

### Run `go run . [PORT]` to start API server on specified port

#### Example: `go run .` to run on default port 8080

## Endpoints

- `/api/artists` - main endpoint that serves all information about all the artists
- `/images/` (or custom) - endpoint to get artists' images

### Environment variables

You can customize API by changing following environment variables:

- `GROUPIE_API_URL` - data source API, default is `https://groupietrackers.herokuapp.com/api`
- `GROUPIE_IMAGES_DIR` - folder to locally store images from API, default is `./images`
- `GROUPIE_IMAGES_URL` - URL to serve locally stored images, default is `/images/`
