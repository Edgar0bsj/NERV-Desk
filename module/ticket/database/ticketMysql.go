package database

import (
	"github.com/edgar0bsj/nerv-desk/module/ticket/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := "root:123456@tcp(localhost:3306)/go_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.TicketModel{}); err != nil {
		return nil, err
	}

	return db, nil
}
