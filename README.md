# task-groupie-tracker-visualizations
grit:lab Åland Islands 2022

---

Authors: [@maximihajlov](https://github.com/maximihajlov), [@healingdrawing](https://github.com/healingdrawing), [@nattikim](https://github.com/nattikim)

Solved during studying in Gritlab coding school on Åland, November 2022

---

## [Task description and audit questions](https://github.com/01-edu/public/tree/master/subjects/groupie-tracker)

---

## Demo: [groupie.mer.pw](https://groupie.mer.pw/)

## Usage

### Docker Compose: `docker compose up`

(multiple container, nginx)

Server will be started on [localhost:80](http://localhost:80/)

---

### Docker: `bash ./build.sh`

(single container, no nginx)

### Development:

You need `go`, `nodejs` and `npm` installed.

- Run `cd backend && go run .` to start backend server
- Run `cd frontend && npm i && npm run dev` to start frontend development server

Check README in `frontend` and `backend` for more details
