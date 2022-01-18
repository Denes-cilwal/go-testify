FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

COPY . /test_web

RUN go get github.com/rubenv/sql-migrate/...

WORKDIR /test_web

RUN go mod download

RUN go get github.com/go-delve/delve/cmd/dlv

CMD sh /test_web/docker/run.sh