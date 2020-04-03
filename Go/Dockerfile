FROM golang:1.8.7-alpine3.6
RUN apk update
RUN apk add git

WORKDIR /go/src/github.com/go-redis/redis/
RUN git clone https://github.com/go-redis/redis.git .
RUN git checkout v6

WORKDIR /go/src/github.com/mackerelio/
RUN git clone https://github.com/mackerelio/go-osstat

WORKDIR /app
COPY . /app

RUN go build main.go

CMD ["./main"]