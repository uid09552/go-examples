package main

import (
	"context"
	"fmt"
	"greetpb"
	"io"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

func main() {
	cc2, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect")
	}
	c2 := greetpb.NewSummationServiceClient(cc2)
	val, _ := Unary2(context.Background(), c2)
	print(val)
	fmt.Println("starting client")
	simpleRequest()
	cc, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect")
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("start server streaming")
	DoServerStreaming(c)
	fmt.Println("Start client streaming")
	DoClientStreaming(c)
	time.Sleep(5000 * time.Millisecond)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		var counter int = 0
		for {
			counter++
			res, _ := Unary(ctx, c)
			fmt.Println(res)
			fmt.Println(counter)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func simpleRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go func() {
		fmt.Println("starting wait overslept")
		select {
		//case <-time.After(5 * time.Second):
		//	fmt.Println("overslept")
		//	cancel()
		case <-ctx.Done():
			fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		}
	}()
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://httpstat.us/200?sleep=4000", nil)
	req = req.WithContext(ctx)
	res, error := client.Do(req)
	print(error)
	print(res)
}

func Unary(ctx context.Context, c greetpb.GreetServiceClient) (string, error) {
	log.Println("Do unary")
	req := &greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: "Max", LastName: "Hh"}}
	res, err := c.Greet(ctx, req)
	if err != nil {
		log.Fatalf("failed", err)
		return "", err
	}
	fmt.Println(res.Result)
	return res.Result, nil
}

func Unary2(ctx context.Context, c greetpb.SummationServiceClient) (int64, error) {
	req := &greetpb.SummationRequest{Summation: &greetpb.Summating{Num1: 1, Num2: 2}}
	res, err := c.Sum(ctx, req)
	if err != nil {
		log.Fatalf("failed", err)
	}
	return res.Result, nil
}

func DoClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("start client streaming..")
	streamClt, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("failed stream ctl", err)
	}
	requests := []string{"hans", "Peter", "Fritz"}
	for _, r := range requests {
		fmt.Printf("sending client ", r)
		streamClt.Send(&greetpb.GreetRequest{Greeting: &greetpb.Greeting{FirstName: r, LastName: "zufall"}})
	}
	resp, err := streamClt.CloseAndRecv()
	if err != nil {

		fmt.Printf("error recv resp", err)
	}
	fmt.Printf("response of stream client", resp)
}

func DoServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("start server streaming client")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Hans",
			LastName:  "Wurst",
		},
	}
	streamRes, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		fmt.Errorf("error:", err)
	}
	for {
		msg, err := streamRes.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("error: ", err)
		}
		fmt.Printf("receive from stream")
		fmt.Println(msg.GetResult())
	}
}
