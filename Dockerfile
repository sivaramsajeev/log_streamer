FROM golang:bullseye AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app/go-streamer . 


FROM alpine:edge

WORKDIR /app/

COPY --from=build /app/go-streamer /app/go-streamer 

CMD ["./go-streamer"]