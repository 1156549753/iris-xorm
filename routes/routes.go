package routes

import (
	"free/controllers"
	"free/services"
	"github.com/kataras/iris"
	// "github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/mvc"
)

func Register(app *iris.Application) {
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
	//	AllowedHeaders:   []string{"*"},
	//	ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	//	AllowCredentials: true,
	//})
	//app.Use(crs)
	userService := services.NewUserService()
	api := mvc.New(app.Party("/api"))
	api.Register(userService)
	api.Handle(new(controllers.UserController))
}
