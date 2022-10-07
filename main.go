package main

import (
	"gin-api/database"
	"gin-api/routers"
	"os"
)

func main() {
	database.StartDB()
	
	var PORT = os.Getenv("PORT")

	routers.StartServer().Run(":" + PORT)
}
