FROM golang:1.24.2-alpine

WORKDIR /app
COPY star_trek_test.go .
COPY go.mod .

RUN go mod tidy 
