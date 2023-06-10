package dao

import (
	"Gone/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB initializes the database connection and migration
func InitDB() *gorm.DB {
	//dsn := "LvKeliang:lkl12190711@tcp(120.79.27.213:3306)/chess?charset=utf8mb4&parseTime=True&loc=Local" // change this according to your own settings
	dsn := "root:123456@tcp(localhost:3306)/chess?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&model.User{}, &model.Board{}, &model.Move{}) // migrate the models to the database tables
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	return db
}
