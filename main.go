package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/pankaj-nupay/ginAndGorm/database"
	"github.com/pankaj-nupay/ginAndGorm/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = database.NewConnection()

	if err != nil {
		log.Fatal("could not load the database")
	}
	routes.Router()
}
