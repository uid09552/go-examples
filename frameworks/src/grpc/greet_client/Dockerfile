FROM golang

ADD ./greet_client/ /go/src/grpc/client
ADD ./greetpb/ /go/src/grpc/greetpb/
WORKDIR /go/src/grpc/client
RUN go get
RUN go build .

ENTRYPOINT ["./client"]

EXPOSE 50052
