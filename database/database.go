package database

import (
	"fmt"
	"log"
	"webapp/go/model"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDB() {
	vi := viper.New()
	vi.SetConfigFile(".env")
	vi.ReadInConfig()


	dsn := fmt.Sprintf("host=localhost user=postgres password=%v dbname=gorm port=5432 sslmode=disable",vi.GetString("DB_PASS"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection failed to the database")
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})

	DB = Dbinstance{
		Db: db,
	}

}