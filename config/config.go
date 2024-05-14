package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_api_poc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
