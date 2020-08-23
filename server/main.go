package main

import (
	"grpc_api_todolist/config"
	"grpc_api_todolist/pkg/v1/api"
	"log"
	"net"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"
)

func main() {
	var db *pg.DB
	grpcServer := grpc.NewServer()
	con := config.NewToDoService(db)
	api.RegisterToDoServiceServer(grpcServer, con)
	l, err := net.Listen("tcp", ":8300")
	log.Println(err)
	grpcServer.Serve(l)
}
