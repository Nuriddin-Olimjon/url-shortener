FROM golang:1.19-alpine3.17
WORKDIR /app

RUN go install github.com/cespare/reflex@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
