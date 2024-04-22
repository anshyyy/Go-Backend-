package main

import (
	"example.com/todoServer/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	routes.InitalizeRoutes(r)
	func() {
		if err := r.Run(":3000"); err != nil {
			log.Fatal("Error occured starting the server", err)
		}
	}()
}
