package main

import (
	"day-7/controllers"
	"day-7/routers"
)

var PORT = "127.0.0.1:8080"

func main() {
	controllers.StartDB()
	routers.StartServer().Run(PORT)
}
