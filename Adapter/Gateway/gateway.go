package gateway

import (
	"database/sql"

	_ "github.com/lib/pq"

	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"

	tasks "github.com/yuya0729/light-clean-architecture/Adapter/Gateway/tasks"
	users "github.com/yuya0729/light-clean-architecture/Adapter/Gateway/users"
)

var DB *sql.DB

// usecaseから呼ばれる
// SQL叩く
// entityに渡す

// users
func GetUsers() ([]*entity.User, error) {
	return users.GetUsers(DB)
}

func GetUser(userID int) (*entity.User, *myerror.MyError) {
	return users.GetUser(DB, userID)
}

// tasks
func GetTasks() ([]*entity.Task, *myerror.MyError) {
	return tasks.GetTasks(DB)
}

func GetTask(userID int, taskID int) (*entity.Task, *myerror.MyError) {
	return tasks.GetTask(DB, userID, taskID)
}

func CreateTask(userID int, title string) *myerror.MyError {
	return tasks.CreateTask(DB, userID, title)
}

func UpdateTask(userID int, title string, taskID int) *myerror.MyError {
	return tasks.UpdateTask(DB, userID, title, taskID)
}

func DeleteTask(userID int, taskID int) *myerror.MyError {
	return tasks.DeleteTask(DB, userID, taskID)
}
