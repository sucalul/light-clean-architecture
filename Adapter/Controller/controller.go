package controller

import (
	"net/http"

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

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	tasks.DeleteTask(w, r)
}
