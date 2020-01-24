package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"url-shortener-api/config"
	"url-shortener-api/database"
	"url-shortener-api/middlewares"
	"url-shortener-api/routes"
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

