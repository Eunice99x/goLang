package database

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	vi := viper.New()
	vi.SetConfigFile(".env")
	vi.ReadInConfig()


	dsn := fmt.Sprintf("host=localhost user=postgres password=%v dbname=gorm port=5432 sslmode=disable",vi.GetString("DB_PASS"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("Database connection successful.")
	return db
}
