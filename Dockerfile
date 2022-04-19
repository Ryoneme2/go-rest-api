FROM golang:1.18.0-alpine3.15 as base

RUN mkdir -p /usr/src/go

WORKDIR /usr/src/go

COPY . .

CMD ["nodemon","--exec", "go", "./src/github.com/Ryoneme2/go-restful-api/cmd/server/main.go","--signal", "SIGTERM"]

FROM cosmtrek/air as air

WORKDIR /usr/src/go

