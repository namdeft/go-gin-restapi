FROM golang:1.21-alpine3.19 AS builder

WORKDIR /app
COPY . .
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
RUN go get github.com/gin-gonic/gin
RUN go build -o main main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .

EXPOSE 8080
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main
