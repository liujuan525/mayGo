FROM golang:alpine

WORKDIR $GOPATH/may-go
COPY . $GOPATH/may-go
RUN go build .

EXPOSE 8008
CMD ["./mayGo"]