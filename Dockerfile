FROM golang:alpine

WORKDIR /opt
ADD ./helloworld.go /opt

RUN apk update && apk upgrade && \
    apk add --no-cache git openssh && \
    /usr/local/go/bin/go get -u github.com/go-redis/redis

ENTRYPOINT ["/usr/local/go/bin/go", "run", "/opt/helloworld.go"]