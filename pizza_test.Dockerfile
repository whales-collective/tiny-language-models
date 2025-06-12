FROM golang:1.24.2-alpine

WORKDIR /app
COPY pizza_test.go .
COPY go.mod .

RUN go mod tidy 
