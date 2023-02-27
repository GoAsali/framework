FROM golang:latest as run

RUN mkdir /app

RUN go install github.com/codegangsta/gin@latest

WORKDIR /app

COPY ./go.mod .
COPY ./go.sum .

RUN go mod tidy


CMD ["gin", "-i", "run","main.go"]