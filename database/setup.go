package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "root:" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
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
