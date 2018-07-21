FROM golang

ADD . /go/src/docker/frontend
WORKDIR /go/src/docker/frontend
RUN go build .

ENTRYPOINT ["./frontend"]

EXPOSE 8080
