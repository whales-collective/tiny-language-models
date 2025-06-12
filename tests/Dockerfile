FROM golang:1.24.2-alpine

WORKDIR /app
COPY main.go .
COPY go.mod .

RUN go mod tidy 
