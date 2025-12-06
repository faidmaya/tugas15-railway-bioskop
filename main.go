package main

import (
	"os"
	"tugas13-bioskop/database"
	"tugas13-bioskop/routers"
)

func main() {
	database.ConnectDB()

	r := routers.StartServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
