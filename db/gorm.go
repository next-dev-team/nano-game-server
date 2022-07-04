package db

import (
	"log"
	"nano-game-server/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/game_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicln("[ORM] err: ", err)
	}
	// auto migration
	db.AutoMigrate(&model.Dice{})
	log.Print("[ORM] Database connection initialized.")
	return db
}
