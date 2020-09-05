package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/nikhil133/grpc_api_todolist/pkg/v1/api"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:local")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to server :", err.Error())
	}
	defer conn.Close()
	c := api.NewToDoServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	/*t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	pfx := t.Format(time.RFC3339Nano)
	req := api.CreateRequest{
		Api: "1",
		/*Todo: &api.ToDo{
			Id:          3,
			Title:       "Workout (" + pfx + ")",
			Description: "Do cross fit (" + pfx + ")",
			Reminder:    reminder,
		},
	}
	res, err := c.Create(ctx, &req)
	log.Println(res, "\nError ", err)*/
	req := api.ReadAllRequest{
		Api: "1",
	}

	resp, err := c.ReadAll(ctx, &req)
	log.Println(resp, "\nError ", err)
}
