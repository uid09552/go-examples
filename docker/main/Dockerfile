FROM golang

ADD . /go/src/docker/main
WORKDIR /go/src/docker/main
RUN go get
RUN go build .

ENTRYPOINT ["./main"]

EXPOSE 8080
