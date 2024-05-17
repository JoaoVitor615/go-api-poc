package config

import (
	"github.com/JoaoVitor615/go-api-poc/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMysql() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	//Create db and connet
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_api_poc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing mysql: %v", err)
		return nil, err
	}

	//Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migration: %v", err)
		return nil, err
	}

	return db, nil
}
