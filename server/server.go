package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"

	. "github.com/dedpanguru/grpc-waitlist/server/set"
	pb "github.com/dedpanguru/grpc-waitlist/waitlist"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

var WaitList Set[string] = Set[string]{}

type Server struct {
	pb.UnimplementedWaitListServer
}

func (s *Server) OptIn(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println(in.GetName(), "wants to join the waitlist")
	WaitList.Add(strings.ToLower(in.GetName()))
	fmt.Println(WaitList)
	return &pb.Response{
		Placement: fmt.Sprintf("%s joined the waitlist! They are number %d.", in.GetName(), WaitList.IndexOf(strings.ToLower(in.GetName()))),
	}, nil
}

func (s *Server) Check(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println(in.GetName(), "wants to check their placement")
	placement := WaitList.IndexOf(strings.ToLower(in.GetName()))
	if placement < 0 {
		return nil, errors.New("you aren't in the waitlist")
	}
	return &pb.Response{
		Placement: fmt.Sprintf("You are number %d.", placement),
	}, nil
}

func (s *Server) OptOut(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println(in.GetName(), "wants to leave the waitlist")
	WaitList.Remove(strings.ToLower(in.GetName()))
	return &pb.Response{
		Placement: "You have been removed from the waitlist",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listed: %v\n", err)
	}
	srvr := grpc.NewServer()
	pb.RegisterWaitListServer(srvr, &Server{})
	log.Println("Server Listening on port", port)
	if err := srvr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
