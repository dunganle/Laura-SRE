FROM golang

COPY main.go .

RUN go get -u github.com/go-redis/redis

RUN go build main.go

EXPOSE 8080

CMD ["./main"]

