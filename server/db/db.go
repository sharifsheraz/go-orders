package db

import (
	"log"
	"os"

	"github.com/sharifsheraz/go-orders/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatalln("Missing database url in env")
	}

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = db.AutoMigrate(&models.Company{}, &models.Customer{}, &models.Order{}, &models.OrderItem{}, &models.Delivery{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	return db
}
