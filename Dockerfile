FROM golang:1.18-alpine
LABEL maintainer=nikolaev.makc@gmail.com
RUN apk add --no-cache git
WORKDIR /go/src/docker-fsnotify-event-handler
COPY . .
RUN go get -d -v ./... && go install -v ./... && rm -rf /go/src
CMD ["docker-fsnotify-event-handler start"]
