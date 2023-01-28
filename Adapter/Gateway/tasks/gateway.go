package tasks

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"
)

func GetTasks(DB *sql.DB) ([]*entity.Task, *myerror.MyError) {
	task := entity.Task{}
	tasks := []*entity.Task{}
	rows, err := DB.Query(`
		SELECT
			tasks.id as id,
			tasks.title,
			users.name
		FROM
			tasks
			JOIN
				users
			ON users.id = tasks.user_id
	`)
	if err != nil {
		log.Println(err)
		return nil, myerror.New(404, err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Title, &task.Name); err != nil {
			return nil, myerror.New(404, err.Error())
		}
		tasks = append(tasks, &entity.Task{ID: task.ID, Title: task.Title, Name: task.Name})
	}
	return tasks, nil
}

func GetTask(c echo.Context, DB *sql.DB, userID int, taskID int) (*entity.Task, *myerror.MyError) {
	task := &entity.Task{}
	err := DB.QueryRow(`
		SELECT
			tasks.id as id,
			tasks.title,
			users.name
		FROM
			tasks
			JOIN
				users
			ON users.id = tasks.user_id
		WHERE tasks.user_id = $1
			AND tasks.id = $2
	`, userID, taskID).Scan(&task.ID, &task.Title, &task.Name)
	if err != nil {
		log.Println(err)
		return nil, myerror.New(404, err.Error())
	}
	return task, nil
}

func CreateTask(c echo.Context, DB *sql.DB, userID int, title string) *myerror.MyError {
	ins, err := DB.Prepare("INSERT INTO tasks (user_id, title) VALUES ($1, $2)")
	if err != nil {
		log.Println(err)
		return myerror.New(404, err.Error())
	}
	ins.Exec(userID, title)
	return nil
}

func UpdateTask(c echo.Context, DB *sql.DB, userID int, title string, taskID int) *myerror.MyError {
	upd, err := DB.Prepare("UPDATE tasks SET user_id = $1, title = $2 WHERE id = $3")
	if err != nil {
		log.Println(err)
		return myerror.New(404, err.Error())
	}
	upd.Exec(userID, title, taskID)
	return nil
}

func DeleteTask(c echo.Context, DB *sql.DB, userID int, taskID int) *myerror.MyError {
	del, err := DB.Prepare("DELETE FROM tasks WHERE id = $1 AND user_id = $2")
	if err != nil {
		log.Println(err)
		return myerror.New(404, err.Error())
	}
	del.Exec(taskID, userID)
	return nil
}
