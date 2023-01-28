package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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

func GetUser(w http.ResponseWriter, r *http.Request) {
	code := 200
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		code = 404
	}
	u, err := interactor.GetUser(userID)
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
