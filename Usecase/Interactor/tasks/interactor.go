package tasks

import (
	"encoding/json"
	"net/http"

	gateway "github.com/yuya0729/light-clean-architecture/Adapter/Gateway"
	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い

func GetTasks() ([]*entity.Task, *myerror.MyError) {
	t, err := gateway.GetTasks()
	if err != nil {
		return nil, err
	}
	return t, nil
}

func BindCreateUpdateTask(r *http.Request) (*entity.CreateTask, *myerror.MyError) {
	task := entity.CreateTask{}
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		return nil, myerror.New(400, err.Error())
	}
	return &task, nil
}

func IsExistsTask(userID int, taskID int) *myerror.MyError {
	if _, err := gateway.GetTask(userID, taskID); err != nil {
		return err
	}
	return nil
}

func CreateTask(userID int, title string) *myerror.MyError {
	if err := gateway.CreateTask(userID, title); err != nil {
		return err
	}
	return nil
}

func UpdateTask(userID int, title string, taskID int) *myerror.MyError {
	if err := gateway.UpdateTask(userID, title, taskID); err != nil {
		return err
	}
	return nil
}

func DeleteTask(userID int, taskID int) *myerror.MyError {
	if err := gateway.DeleteTask(userID, taskID); err != nil {
		return err
	}
	return nil
}
