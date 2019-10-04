FROM golang

ADD . /go/src/greet/server
WORKDIR /go/src/greet/server
RUN go get
RUN go build .

ENTRYPOINT ["./server"]

EXPOSE 50051
