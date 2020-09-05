package main

import (
	"log"
	"net"

	"github.com/nikhil133/grpc_api_todolist/pkg/v1/api"
	"github.com/nikhil133/grpc_api_todolist/pkg/v1/services"

	"github.com/go-pg/pg"
	"google.golang.org/grpc"
)

func main() {
	var db *pg.DB
	grpcServer := grpc.NewServer()
	con := services.NewToDoService(db)
	api.RegisterToDoServiceServer(grpcServer, con)
	l, err := net.Listen("tcp", ":8300")
	log.Println(err)
	log.Println("Go server going to start at 8300")
	grpcServer.Serve(l)

}
