package routes

import (
	"github.com/VishwasShashidhar/car-os/src/controllers"
	"github.com/kataras/iris/v12"
)

// SetupHomeRoutes ... Sets up the home routes
func SetupHomeRoutes(app *iris.Application) {
	homeAPI := app.Party("/home")
	{
		homeAPI.Use(iris.Compression)
		homeAPI.Get("/", controllers.GreetUser)
	}

}
