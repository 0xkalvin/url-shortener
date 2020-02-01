package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"url-shortener-api/src/config"
	"url-shortener-api/src/database"
	"url-shortener-api/src/middlewares"
	"url-shortener-api/src/routes"
)


func main() {

	config.LoadConfig()
	db := database.SetupDatabase()

	app := gin.Default()
	
	app.Use(database.Inject(db))
	app.Use(cors.Default())
	
	routes.SetupRoutes(app)
	app.NoRoute(middlewares.NotFoundHandler)

	app.Run() 
}

