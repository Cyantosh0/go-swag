package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"Cyantosh0/go-swag/config"
	_ "Cyantosh0/go-swag/docs"
	"Cyantosh0/go-swag/model"
	"Cyantosh0/go-swag/route"
)

// @title User API documentation
// @version 1.0.0

// @host localhost:5000
// @BasePath /user

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.DB = config.SetupDatabase()
	config.DB.AutoMigrate(&model.User{})

	r := route.SetupRouter()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + os.Getenv("SERVER_PORT"))
}
