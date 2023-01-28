package driver

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	controller "github.com/yuya0729/light-clean-architecture/Adapter/Controller"
	gateway "github.com/yuya0729/light-clean-architecture/Adapter/Gateway"
)

// routerとかdb接続とか
func Serve() {
	// db接続
	var err error
	dsn := "user=postgres host=postgres port=5432 dbname=postgres password=postgres sslmode=disable"
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	// gateway DBにDBを渡す
	gateway.DB = DB

	// APIルーティング
	// api.DELETE("/tasks/:id", controller.DeleteTask)
	// create user

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/users", controller.GetUsers)
		r.Get("/users/{userID}", controller.GetUser)
		r.Get("/tasks", controller.GetTasks)
		r.Post("/tasks", controller.CreateTask)
		r.Put("/tasks/{taskID}", controller.UpdateTask)
	})
	http.ListenAndServe(":8080", r)
}
