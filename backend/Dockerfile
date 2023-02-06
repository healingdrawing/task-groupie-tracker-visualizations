FROM golang:alpine AS builder

WORKDIR /app/
COPY . .

RUN go build .

FROM alpine

# RUN apk add curl

COPY --from=builder /app/groupie-tracker /app/bin/
CMD /app/bin/groupie-tracker

# HEALTHCHECK --interval=3s CMD curl --fail http://localhost:8080/api/artists || exit 1

EXPOSE 8080