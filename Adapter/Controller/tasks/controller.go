package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	interactor "github.com/yuya0729/light-clean-architecture/Usecase/Interactor"
)

// TODO:
// interface導入？

// TODO:
// controllerでやっているバリデーションを各層に分散
// ref: https://qiita.com/nakaakist/items/11195ac5290450fbc9f9

// TODO:
// errorの体系化

func GetTasks(w http.ResponseWriter, r *http.Request) {
	code := 200
	t, err := interactor.GetTasks()
	if err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(code)

	// httpResponseの内容を書き込む
	buf, _ := json.MarshalIndent(t, "", "    ")
	_, _ = w.Write(buf)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	code := 200
	task, err := interactor.BindCreateUpdateTask(r)
	if err != nil {
		switch err.Code {
		case 400:
			code = 400
		}
	}
	if err := interactor.IsExistsUser(task.UserID); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	if err = interactor.CreateTask(task.UserID, task.Title); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(code)

	// httpResponseの内容を書き込む
	// TODO: responseちゃんと返す
	buf, _ := json.MarshalIndent("", "", "    ")
	_, _ = w.Write(buf)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	code := 200
	taskID, e := strconv.Atoi(chi.URLParam(r, "taskID"))
	if e != nil {
		code = 400
	}
	task, err := interactor.BindCreateUpdateTask(r)
	if err != nil {
		switch err.Code {
		case 400:
			code = 400
		}
	}
	if err = interactor.IsExistsTask(task.UserID, taskID); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	if err = interactor.UpdateTask(task.UserID, task.Title, taskID); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(code)

	// httpResponseの内容を書き込む
	// TODO: responseちゃんと返す
	buf, _ := json.MarshalIndent("", "", "    ")
	_, _ = w.Write(buf)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	code := 200
	taskID, e := strconv.Atoi(chi.URLParam(r, "taskID"))
	if e != nil {
		code = 400
	}
	userID, e := strconv.Atoi(chi.URLParam(r, "user_id"))
	if e != nil {
		code = 400
	}
	if err := interactor.IsExistsTask(userID, taskID); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	if err := interactor.DeleteTask(userID, taskID); err != nil {
		switch err.Code {
		case 404:
			code = 404
		}
	}
	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(code)

	// httpResponseの内容を書き込む
	// TODO: responseちゃんと返す
	buf, _ := json.MarshalIndent("", "", "    ")
	_, _ = w.Write(buf)
}
