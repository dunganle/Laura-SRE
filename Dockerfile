FROM golang:alpine

COPY main.go .

RUN apk add git

RUN go get -u github.com/go-redis/redis

RUN go build main.go

EXPOSE 9001

CMD ["./main"]

