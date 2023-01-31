package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/osaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, table := range RegisterTables() {
		err = db.Debug().AutoMigrate(table.Table)

		if err != nil {
			log.Fatal(err)
		}
	}

	return db
}
