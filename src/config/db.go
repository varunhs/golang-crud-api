package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := os.Getenv("DB_URL") // To read it from the .env file.

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect to the DB", err)
	}

	sqlDB, err := db.DB()

	if err !=nil {
		log.Fatal("❌ Error getting DB instance:", err)
	}

	if err = sqlDB.Ping(); err  != nil {
		log.Fatal("❌ Database not reachable:", err)
	}

	DB = db
	log.Println("✅ Database Connected")
}
