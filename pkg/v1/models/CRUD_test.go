/*con := c.Config{}
con.InitDB()
log.Println(con.DB)
s := con.Serve
var req v1.ReadAllRequest
req.Api = "1"
todoList, err := models.ReadAllToDo(con, &req)
log.Println("Read all \n", todoList, "\nerror ", err)

var req1 v1.ReadRequest
req1.Api = "1"
req1.Id = 2
todo, err := models.ReadToDo(con, &req1)
log.Println("Read \n", todo, "\nerror ", err)

var req2 v1.UpdateRequest
req2.Api = "1"
t := time.Now().In(time.UTC)
uts, _ := ptypes.TimestampProto(t)
upTodo := v1.ToDo{
	Id:          2,
	Title:       "Docker, Kubernetes and CI/CD",
	Description: "learn containizers and orchastration of containers from udemy with CI/CD pipeline tarvis",
	Reminder:    uts,
}
req2.Todo = &upTodo
updated, _ := models.UpdateToDo(con, &req2)
log.Println("Updated rows ", updated.Update)

var req3 v1.DeleteRequest
req3.Api = "1"
req3.Id = 2
deleted, _ := models.DeleteToDo(con, &req3)
log.Println("Deleted rows ", deleted.Delete)
//if err != nil {
//
//} else {
//	log.Fatal(err)
//}*/