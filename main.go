package main

import (
	"Cyantosh0/go-swag/config"
	"Cyantosh0/go-swag/model"
	"Cyantosh0/go-swag/route"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := config.SetupDatabase()
	db.AutoMigrate(&model.User{})

	r := route.SetupRouter()
	r.Run()
}
