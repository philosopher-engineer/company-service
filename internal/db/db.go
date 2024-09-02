package db

import (
	"company-service/internal/company"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func InitDB() *gorm.DB {
	config := viper.New()
	config.SetConfigFile(".env")
	config.ReadInConfig()

	db, err := gorm.Open("postgres", config.GetString("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	db.LogMode(true)
	db.AutoMigrate(&company.Company{})

	return db
}
