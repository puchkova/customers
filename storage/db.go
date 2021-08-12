package storage

import (
	"customers/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
