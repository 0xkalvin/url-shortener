package main

import (
	"github.com/0xkalvin/url-shortener/database"
	"github.com/0xkalvin/url-shortener/server"
)

func main() {

	database.InitializeDynamoDB()
	database.InitializeRedis()

	server.Run()
}
