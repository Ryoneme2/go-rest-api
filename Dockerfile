FROM golang:1.18.0-alpine3.15 as base

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

CMD ["go", "run", "./src/github.com/Ryoneme2/go-restful-api/cmd/server/main.go"]