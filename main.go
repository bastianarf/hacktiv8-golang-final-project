package main

import (
	"example.id/mygram/database"
	_ "example.id/mygram/docs"
	"example.id/mygram/routers"
)

func main() {
	database.StartDB()
	var port = ":8080"
	routers.StartServer().Run(port)
}
