FROM golang:1.11-alpine
MAINTAINER nikolaev.makc@gmail.com
RUN apk add --no-cache git
WORKDIR /go/src/docker-fsnotify-event-handler
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["docker-fsnotify-event-handler"]
