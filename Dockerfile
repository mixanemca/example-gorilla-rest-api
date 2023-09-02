##
## Build
##
FROM golang:1.21.0 AS build

WORKDIR /app/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build
