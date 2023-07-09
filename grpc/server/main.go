package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	model "github.com/faaizz/learnings/bengineering/grpc/server/model"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ts := model.TodoServerImpl{}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	model.RegisterTodoServer(s, &ts)
	s.Serve(lis)
}
