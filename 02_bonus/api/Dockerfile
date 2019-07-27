FROM golang:1.10

WORKDIR /go/src/api

COPY . .
COPY src/Users /go/src/Users
COPY src/Redis /go/src/Redis

ENV GO111MODULE=on

RUN go get github.com/go-redis/redis
RUN go get github.com/gorilla/mux
RUN go install Users
RUN go install Redis

CMD ["go", "run", "main.go", "router.go"]
