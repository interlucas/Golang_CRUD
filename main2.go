package main

import (
	"log"
	"github.com/interlucas/Golang_CRUD/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file \n",
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}