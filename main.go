package main

import (
	"fmt"
	"log"
	"os"
	"github.com/interlucas/Golang_CRUD/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))
}
