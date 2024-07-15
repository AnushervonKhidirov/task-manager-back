package database

import (
	"main/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := "main:main@tcp(127.0.0.1:3306)/task_manager?charset=utf8mb4&parseTime=True"
	db, connectionErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if connectionErr != nil {
		return connectionErr
	}

	db.AutoMigrate(&model.User{})

	DB = db

	return nil
}

func AddUser(user *model.User) model.User {
	var result model.User

	DB.Create(&user)
	DB.Last(&result)

	return result
}
