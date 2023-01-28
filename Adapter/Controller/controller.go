package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	tasks "github.com/yuya0729/light-clean-architecture/Adapter/Controller/tasks"
	users "github.com/yuya0729/light-clean-architecture/Adapter/Controller/users"
)

// driverで定義されたエンドポイントの関数を定義する

// users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users.GetUsers(w, r)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	users.GetUser(w, r)
}

// tasks
func GetTasks(c echo.Context) error {
	return tasks.GetTasks(c)
}

func CreateTask(c echo.Context) error {
	return tasks.CreateTask(c)
}

func UpdateTask(c echo.Context) error {
	return tasks.UpdateTask(c)
}

func DeleteTask(c echo.Context) error {
	return tasks.DeleteTask(c)
}
