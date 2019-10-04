FROM golang

ADD ./greet_server/ /go/src/grpc/server
ADD ./greetpb/ /go/src/grpc/greetpb/
WORKDIR /go/src/grpc/server
RUN go get
RUN go build .

ENTRYPOINT ["./server"]

EXPOSE 50051
