package interactor

import (
	"net/http"

	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"

	tasks "github.com/yuya0729/light-clean-architecture/Usecase/Interactor/tasks"
	users "github.com/yuya0729/light-clean-architecture/Usecase/Interactor/users"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetUsers() ([]*entity.User, error) {
	return users.GetUsers()
}

func GetUser(userID int) (*entity.User, error) {
	return users.GetUser(userID)
}

func IsExistsUser(userID int) *myerror.MyError {
	return users.IsExistsUser(userID)
}

func GetTasks() ([]*entity.Task, *myerror.MyError) {
	return tasks.GetTasks()
}

func BindCreateUpdateTask(r *http.Request) (*entity.CreateTask, *myerror.MyError) {
	return tasks.BindCreateUpdateTask(r)
}

func CreateTask(userID int, title string) *myerror.MyError {
	return tasks.CreateTask(userID, title)
}

func UpdateTask(userID int, title string, taskID int) *myerror.MyError {
	return tasks.UpdateTask(userID, title, taskID)
}

func IsExistsTask(userID int, taskID int) *myerror.MyError {
	return tasks.IsExistsTask(userID, taskID)
}

func DeleteTask(userID int, taskID int) *myerror.MyError {
	return tasks.DeleteTask(userID, taskID)
}
