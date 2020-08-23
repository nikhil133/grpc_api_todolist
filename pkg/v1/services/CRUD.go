package services

import (
	"context"
	"time"

	"github.com/nikhil133/grpc_api_todolist/pkg/v1/api"

	"github.com/go-pg/pg"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	apiVer = "1"
)

//ToDoService is the implementation of grpc service structure
type ToDoService struct {
	DB *pg.DB
}

//NewToDoService constructor for ToDoService
func NewToDoService(db *pg.DB) *ToDoService {
	return &ToDoService{DB: InitDB()}
}

//InitDB connecting postgres sql
func InitDB() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "bon",
		Password: "",
		Database: "todo",
	})

}

//CheckApi validate api version
func CheckApi(api string) error {
	if apiVer != api {
		return status.Errorf(codes.Unimplemented,
			"unsupported API version: service implements API version '%s', but asked for '%s'", apiVer, api)
	}

	return nil
}

//ToDoList struct for psql table todo_list
type ToDoList struct {
	TableName   struct{}  `sql:"todo_list"`
	ID          int64     `sql:"id"`
	Title       string    `sql:"title"`
	Description string    `sql:"description"`
	Reminder    time.Time `sql:"reminder"`
}

//Create : Insert query on todo_list table
func (c *ToDoService) Create(ctx context.Context, req *api.CreateRequest) (*api.CreateResponse, error) {
	err := CheckApi(req.Api)
	if err != nil {
		return nil, err
	}
	//c := &config.Config{}
	//c.InitDB()
	//defer c.DB.Close()

	var row ToDoList
	if req.Todo.Id != 0 {
		row.ID = req.Todo.Id
	}

	row.Title = req.Todo.Title
	row.Description = req.Todo.Description

	row.Reminder, _ = ptypes.Timestamp(req.Todo.Reminder)
	err = c.DB.Insert(&row)
	var resp api.CreateResponse
	resp.Api = req.Api
	resp.Id = row.ID
	return &resp, err
}

//ReadAll : query to read all content of todo_list table
func (c *ToDoService) ReadAll(ctx context.Context, req *api.ReadAllRequest) (*api.ReadAllResponse, error) {

	err := CheckApi(req.Api)
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

//Read : read by id query on todo_list table
func (c *ToDoService) Read(ctx context.Context, req *api.ReadRequest) (*api.ReadResponse, error) {

	err := CheckApi(req.Api)
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

//Update : update by id query on todo_list table
func (c *ToDoService) Update(ctx context.Context, req *api.UpdateRequest) (*api.UpdateResponse, error) {

	err := CheckApi(req.Api)
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

//Delete : delete by id query on table todo_list table
func (c *ToDoService) Delete(ctx context.Context, req *api.DeleteRequest) (*api.DeleteResponse, error) {

	err := CheckApi(req.Api)
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
