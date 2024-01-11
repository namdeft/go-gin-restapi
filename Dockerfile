FROM golang:1.21-alpine3.19 AS builder

RUN mkdir /app
WORKDIR /app
COPY . /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin

ENTRYPOINT ["CompileDaemon", "--build='go build main.go'", "--command=./main"]