package database

import (
	"fmt"
	"go-start/pkg/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=your_user password=your_password dbname=your_dbname port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %v", err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Cart{}, &model.CartItem{})
	if err != nil {
		return nil, fmt.Errorf("ошибка при миграции: %v", err)
	}

	return db, nil
}
