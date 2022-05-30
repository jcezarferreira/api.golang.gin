package main

import (
	"github.com/jcezarferreira/api.gin/database"
	"github.com/jcezarferreira/api.gin/routes"
)

func main() {
	database.ConnectDB()

	routes.HandleRequests()
}
