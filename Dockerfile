FROM golang:alpine

WORKDIR /app/

COPY ./backend /app/backend
COPY ./frontend /app/frontend

WORKDIR /app/backend
RUN go build -o="/app/" .


FROM node:alpine

COPY --from=0 /app /app

WORKDIR /app/frontend

RUN npm install

EXPOSE 8080 3000

CMD /app/groupie-tracker & (npm run build && npm run start)

