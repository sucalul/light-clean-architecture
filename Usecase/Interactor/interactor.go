package interactor

import (
	"github.com/labstack/echo/v4"
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

func BindCreateUpdateTask(c echo.Context) (*entity.CreateTask, *myerror.MyError) {
	return tasks.BindCreateUpdateTask(c)
}

func CreateTask(c echo.Context, userID int, title string) *myerror.MyError {
	return tasks.CreateTask(c, userID, title)
}

func UpdateTask(c echo.Context, userID int, title string, taskID int) *myerror.MyError {
	return tasks.UpdateTask(c, userID, title, taskID)
}

func IsExistsTask(c echo.Context, userID int, taskID int) *myerror.MyError {
	return tasks.IsExistsTask(c, userID, taskID)
}

func DeleteTask(c echo.Context, userID int, taskID int) *myerror.MyError {
	return tasks.DeleteTask(c, userID, taskID)
}
