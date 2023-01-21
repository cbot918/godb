package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "user/pb/users"
	u "user/util"

	"google.golang.org/grpc"
)

const (
	port = ":50055"
	name = "yale"
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginReply, error) {
	log.Printf("Received Email: %v", in.GetEmail())
	log.Printf("Received Password: %v", in.GetPassword())
	return &pb.LoginReply{Name: name}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	u.Checke(err, "net listen failed")

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})
	fmt.Printf("server listening: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
