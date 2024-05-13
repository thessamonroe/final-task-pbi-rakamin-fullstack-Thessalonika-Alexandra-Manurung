package database

import (
	"fmt"
	"fullStack/app/config"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	db, err := gorm.Open("mysql", dbConfig)
	if err != nil {
		panic("Koneksi gagal")
	}

	DB = db

	migrateModels()
}

func migrateModels() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
