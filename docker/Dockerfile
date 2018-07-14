FROM golang

ADD . /go/src/docker
WORKDIR /go/src/docker
RUN go build .

ENTRYPOINT ["./docker"]

EXPOSE 8080
