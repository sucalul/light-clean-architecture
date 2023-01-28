package users

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	interactor "github.com/yuya0729/light-clean-architecture/Usecase/Interactor"
)

// driverで定義されたエンドポイントの関数を定義する
func GetUsers(w http.ResponseWriter, r *http.Request) {
	code := 200
	u, err := interactor.GetUsers()
	if err != nil {
		// TODO:ちゃんとエラー書く
		code = 404
	}

	// レスポンスヘッダーの設定
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// ステータスコードを設定
	w.WriteHeader(code)

	// httpResponseの内容を書き込む
	buf, _ := json.MarshalIndent(u, "", "    ")
	_, _ = w.Write(buf)
}

func GetUser(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	u, err := interactor.GetUser(c, userID)
	if err != nil {
		msg := fmt.Sprintf(`{"message": %s`, err)
		return c.JSON(http.StatusBadRequest, msg)
	}
	return c.JSON(http.StatusOK, u)
}
