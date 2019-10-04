package main

import (
	"context"
	"fmt"
	"grpc/greetpb"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

//unary server
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{Result: result}
	return res, nil
}

//unary server
func (*server) Sum(ctx context.Context, req *greetpb.SummationRequest) (*greetpb.SummationResponse, error) {
	res := req.GetSummation().GetNum1() + req.GetSummation().GetNum2()
	return &greetpb.SummationResponse{Result: res}, nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	names := []string{}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("finished client stream")
			result := "Hello " + strings.Join(names, " ")
			return stream.SendAndClose(&greetpb.GreetLongResponse{
				Result: result,
			})
		}
		if err != nil {
			fmt.Errorf("error reading client stream", err)
		}
		names = append(names, msg.GetGreeting().GetFirstName())
	}

	return nil
}

//stream
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	fmt.Println("Greet function was invoked")
	for i := 0; i < 10; i++ {
		response := &greetpb.GreetManyTimesResponse{
			Result: "hello " + firstName,
		}
		stream.Send(response)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

func main() {
	log.Printf("server is starting")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	s := grpc.NewServer()
	server := &server{}
	greetpb.RegisterGreetServiceServer(s, server)
	greetpb.RegisterSummationServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve")
	}

}
