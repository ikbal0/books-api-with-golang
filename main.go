package main

import (
	"gin-api/database"
	"gin-api/routers"
)

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
	database.StartDB()
}
