package main

import (
	"context"
	"log"
	"time"
	pb "user/pb/users"

	u "user/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr     = "localhost:50055"
	email    = "yale918@gmail.com"
	password = "12345"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	u.Checke(err, "grpc dial error")
	defer conn.Close()
	c := pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Login(ctx, &pb.LoginRequest{Email: email, Password: password})
	u.Checke(err, "pb client Login failed")
	// u.Logg(r)
	log.Fatalf("Server Res: %s", r.GetName())

}
