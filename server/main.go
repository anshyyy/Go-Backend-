package main

import (
	"example.com/todoServer/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	routes.InitalizeRoutes(r)
	if err := godotenv.Load(); err != nil {
		log.Panic("Error loading .env file")
	}
	port := os.Getenv("PORT")
	func() {
		if err := r.Run(":" + port); err != nil {
			log.Fatal("Error occured starting the server", err)
		}
	}()
}
