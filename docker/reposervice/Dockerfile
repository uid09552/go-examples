FROM golang

ADD . /go/src/docker/reposervice
WORKDIR /go/src/docker/reposervice
RUN go get
RUN go build .

ENTRYPOINT ["./reposervice"]

EXPOSE 8085
