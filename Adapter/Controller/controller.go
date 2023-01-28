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
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks.GetTasks(w, r)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	tasks.CreateTask(w, r)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	tasks.UpdateTask(w, r)
}

func DeleteTask(c echo.Context) error {
	return tasks.DeleteTask(c)
}
