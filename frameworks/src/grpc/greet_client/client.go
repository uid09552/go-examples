package main

import (
	"context"
	"fmt"
	"grpc/greetpb"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting client")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50052", nil)
	fmt.Println("end client")
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := invokeAll()
	fmt.Fprintf(w, "invoked grpc result: %s", result)
}

func invokeAll() string {
	cc2, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect")
	}
	c2 := greetpb.NewSummationServiceClient(cc2)
	val, _ := Unary2(context.Background(), c2)
	print(val)
	fmt.Println("starting client")
	simpleRequest()
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect")
	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Println("start server streaming")
	DoServerStreaming(c)
	fmt.Println("Start client streaming")
	DoClientStreaming(c)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result := []string{}
	go func() {
		var counter int = 0
		for {
			counter++
			res, err := Unary(ctx, c)
			fmt.Println(res)
			fmt.Println(counter)
			result = append(result, res)
			if err != nil {
				break
			}
		}
	}()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("after timeout canceled")
		cancel()
	case <-ctx.Done():
		log.Fatalln("canceled")
	}
	return strings.Join(result, " ")
}

func simpleRequest() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
	req, _ := http.NewRequest("GET", "http://httpstat.us/200?sleep=1000", nil)
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
		log.Printf("unary failed", err)
		return "", err
	}
	fmt.Println(res.Result)
	return res.Result, nil
}

func Unary2(ctx context.Context, c greetpb.SummationServiceClient) (int64, error) {
	req := &greetpb.SummationRequest{Summation: &greetpb.Summating{Num1: 1, Num2: 2}}
	res, err := c.Sum(ctx, req)
	if err != nil {
		log.Printf("unary 2 failed", err)
		return 0, err
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
