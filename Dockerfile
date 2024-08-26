FROM golang:bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/go-streamer . 


FROM ubuntu:jammy

WORKDIR /app/

COPY --from=build /app/go-streamer /app/go-streamer 

ENTRYPOINT ["./go-streamer"]