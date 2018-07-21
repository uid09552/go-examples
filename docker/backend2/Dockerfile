FROM golang

ADD . /go/src/docker/backend
WORKDIR /go/src/docker/backend
RUN go build .

ENTRYPOINT ["./backend"]

EXPOSE 8081
