package interactor

import (
	"github.com/labstack/echo/v4"
	entity "github.com/yuya0729/light-clean-architecture/cmd/app/Entity"

	users "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/Interactor/users"
	tasks "github.com/yuya0729/light-clean-architecture/cmd/app/Usecase/tasks"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetUsers(c echo.Context) ([]*entity.User, error) {
	return users.GetUsers(c)
}

func GetUser(c echo.Context, userID int) (*entity.User, error) {
	return users.GetUser(c, userID)
}

func GetTasks(c echo.Context) ([]*entity.Task, error) {
	return tasks.GetTasks(c)
}

func BindCreateTask(c echo.Context) (*entity.CreateTask, error) {
	return tasks.BindCreateTask(c)
}

func IsExistsUser(c echo.Context, userID int) error {
	return tasks.IsExistsUser(c, userID)
}

func CreateTask(c echo.Context, userID int, title string) error {
	return tasks.CreateTask(c, userID, title)
}
