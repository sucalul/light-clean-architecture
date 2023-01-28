package users

import (
	"github.com/labstack/echo/v4"
	gateway "github.com/yuya0729/light-clean-architecture/Adapter/Gateway"
	myerror "github.com/yuya0729/light-clean-architecture/Driver/error"
	entity "github.com/yuya0729/light-clean-architecture/Entity"
)

// controllerから関数を呼び出す
// その関数ではrepositoryを呼び出す
// service的な役割
// interfaceがあっても良い
// memo

func GetUsers() ([]*entity.User, error) {
	u, err := gateway.GetUsers()
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(c echo.Context, userID int) (*entity.User, error) {
	u, err := gateway.GetUser(c, userID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func IsExistsUser(c echo.Context, userID int) *myerror.MyError {
	if _, err := gateway.GetUser(c, userID); err != nil {
		return err
	}
	return nil
}
