package main

import (
	"context"
	"log"
	"time"

	pb "github.com/dedpanguru/grpc-waitlist/waitlist"
	"google.golang.org/grpc"
)

const serverAddress = "localhost:8080"

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect to server: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewWaitListClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var suckers = []string{"Alice", "bob", "carl"}
	for _, name := range suckers {
		response, err := client.OptIn(ctx, &pb.Request{Name: name})
		if err != nil {
			log.Fatalf("there was a problem with opting in: %v\n", err)
		}
		log.Println(response.GetPlacement())
	}
	for _, name := range suckers {
		response, err := client.Check(ctx, &pb.Request{Name: name})
		if err != nil {
			log.Fatalf("there was a problem with checking: %v\n", err)
		}
		log.Println(response.GetPlacement())
	}

}
