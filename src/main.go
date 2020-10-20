package main

import (
	"fmt"
	"github.com/VishwasShashidhar/car-os/src/common"
	_ "github.com/VishwasShashidhar/car-os/src/docs"
	"github.com/VishwasShashidhar/car-os/src/routes"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"log"
)

// @title Car OS
// @version 1.0
// @description This is a simple car management app
// @termsOfService http://swagger.io/terms/

// @contact.name Vishwas Shashidhar
// @contact.url http://www.adwitiya.io/
// @contact.email vishwas@adwitiya.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	common.Init()
	common.LoadConfiguration()

	// Setup Gorm DB Configuration
	common.ConnectDatabase()

	// Setup IRIS & Swagger
	app := iris.New()
	app.Get("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
	routes.SetupHomeRoutes(app)
	port := common.GetConfigurationForKey("APP_PORT")
	log.Printf("Open Swagger on http://localhost:%s/swagger/index.html", port)
	app.Listen(fmt.Sprintf("localhost:%s", port))
}
