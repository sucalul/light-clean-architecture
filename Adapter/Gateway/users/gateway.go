package users

import (
	"database/sql"
	"log"

	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"
)

func GetUsers(DB *sql.DB) ([]*entity.User, error) {
	user := entity.User{}
	users := []*entity.User{}
	rows, err := DB.Query("SELECT id, name FROM users")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, &entity.User{ID: user.ID, Name: user.Name})
	}
	return users, nil
}

func GetUser(DB *sql.DB, userID int) (*entity.User, *myerror.MyError) {
	user := &entity.User{}
	err := DB.QueryRow("SELECT id, name FROM users WHERE id = $1", userID).Scan(&user.ID, &user.Name)
	if err != nil {
		log.Println("err")
		msg := myerror.New(404, err.Error())
		return nil, msg
	}
	return user, nil
}
