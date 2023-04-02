package main

import (
	"books-api/routers"
)

var PORT = ":8080"

func main() {
	routers.StartServer().Run(PORT)
}
