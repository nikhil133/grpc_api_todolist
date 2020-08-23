package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/go-pg/pg"
	"github.com/golang/protobuf/ptypes"
	"github.com/nikhil133/grpc_api_todolist/pkg/v1/api"
	"github.com/nikhil133/grpc_api_todolist/pkg/v1/services"
)

func TestServiceCreate(t *testing.T) {
	var db *pg.DB
	c := services.NewToDoService(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tym := time.Now().In(time.UTC)
	uts, _ := ptypes.TimestampProto(tym)
	req := &api.CreateRequest{
		Api: "1",
		Todo: &api.ToDo{
			Id:          2,
			Title:       "Angular Js",
			Description: "Learn angular js",
			Reminder:    uts,
		},
	}

	_, err := c.Create(ctx, req)
	require.NoError(t, err)
}

func TestServiceRead(t *testing.T) {
	var db *pg.DB
	c := services.NewToDoService(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.ReadRequest{
		Api: "1",
		Id:  2,
	}
	_, err := c.Read(ctx, req)
	require.NoError(t, err)
}

func TestServiceReadAll(t *testing.T) {
	var db *pg.DB
	c := services.NewToDoService(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.ReadAllRequest{
		Api: "1",
	}
	_, err := c.ReadAll(ctx, req)
	require.NoError(t, err)
}

func TestServiceUpdate(t *testing.T) {
	var db *pg.DB
	c := services.NewToDoService(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	tym := time.Now().In(time.UTC)
	uts, _ := ptypes.TimestampProto(tym)

	req := &api.UpdateRequest{
		Api: "1",
		Todo: &api.ToDo{
			Id:          2,
			Title:       "Angular Js",
			Description: "Learn angular js. Completed till DI Services",
			Reminder:    uts,
		},
	}
	_, err := c.Update(ctx, req)
	require.NoError(t, err)
}

func TestServiceDelete(t *testing.T) {
	var db *pg.DB
	c := services.NewToDoService(db)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &api.DeleteRequest{
		Api: "1",
		Id:  2,
	}
	_, err := c.Delete(ctx, req)
	require.NoError(t, err)
}
