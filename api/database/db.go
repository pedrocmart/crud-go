package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pedrocmart/crud-go/api/database/config"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", config.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
