package main

import (
	"Cyantosh0/go-swag/config"
	"Cyantosh0/go-swag/route"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer db.Close()

	// db.AutoMigrate(&model.User{})

	r := route.SetupRouter()
	r.Run()
}
