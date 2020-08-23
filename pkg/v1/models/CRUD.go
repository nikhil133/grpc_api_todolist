package models

import (
	"context"
	"grpc_api_todolist/config"
	"grpc_api_todolist/pkg/v1/api"
	"time"

	"github.com/golang/protobuf/ptypes"
)

type ToDoList struct {
	TableName   struct{}  `sql:"todo_list"`
	ID          int64     `sql:"id"`
	Title       string    `sql:"title"`
	Description string    `sql:"description"`
	Reminder    time.Time `sql:"reminder"`
}

func (c *config.ToDoService) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	err := config.CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var row ToDoList
	row.ID = req.Todo.Id
	row.Title = req.Todo.Title
	row.Description = req.Todo.Description

	row.Reminder, _ = ptypes.Timestamp(req.Todo.Reminder)
	err = c.DB.Insert(&row)
	var resp api.CreateResponse
	resp.Api = req.Api
	resp.Id = row.Id
	return &resp, err
}
func (c *config.ToDoService) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {

	err := config.CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var rows []ToDoList

	err = c.DB.ModelContext(ctx, &ToDoList{}).Select(&rows)

	var resp api.ReadAllResponse
	resp.Api = req.Api
	for i := range rows {
		var todo api.ToDo
		todo.Id = rows[i].ID
		todo.Title = rows[i].Title
		todo.Description = rows[i].Description
		t, _ := ptypes.TimestampProto(rows[i].Reminder)
		todo.Reminder = t
		resp.Todos = append(resp.Todos, &todo)
	}
	return &resp, err

}

func (c *config.ToDoService) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {

	err := config.CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var row ToDoList

	err = c.DB.ModelContext(ctx, &ToDoList{}).Where("id=?", req.Id).Select(&row)
	var resp api.ReadResponse
	var todo api.ToDo
	resp.Api = req.Api
	todo.Id = row.ID
	todo.Title = row.Title
	todo.Description = row.Description
	todo.Reminder, _ = ptypes.TimestampProto(row.Reminder)
	resp.Todo = &todo
	return &resp, err
}

func (c *config.ToDoService) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {

	err := config.CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var row ToDoList
	t, _ := ptypes.Timestamp(req.Todo.Reminder)
	updates, err := c.DB.Model(&ToDoList{}).
		Set("title=?", req.Todo.Title).
		Set("description=?", req.Todo.Description).
		Set("reminder=?", t).
		Where("id=?", req.Todo.Id).
		Update(&row)
	var resp api.UpdateResponse
	resp.Api = req.Api
	resp.Update = int64(updates.RowsAffected())

	return &resp, err

}

func (c *config.ToDoService) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {

	err := config.CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var row ToDoList
	deleted, err := c.DB.Model(&ToDoList{}).
		Where("id=?", req.Id).
		Delete(&row)
	var resp api.DeleteResponse
	resp.Api = req.Api
	resp.Delete = int64(deleted.RowsAffected())
	return &resp, err
}
