package main

import (
	"fmt"
	"log"
	"net"
	pb "user/pb/users"
	u "user/util"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServer
}

const (
	port = ":50055"
)

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
