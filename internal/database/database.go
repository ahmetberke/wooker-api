package database

import (
	"fmt"
	"github.com/ahmetberke/wooker-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type  DBConfig struct {
	Host string
	Port string
	Name string
	User string
	Password string
	SSLMode string
}

func ConnectToDB(config *DBConfig) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", config.Host, config.Port, config.Name, config.User, config.Password, config.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return  nil, err
	}
	log.Printf("Connect to database, on %s:%s", config.Host, config.Port)
	migrations(db)
	afterConnectionWorks(db)
	return db, err
}

func migrations(db *gorm.DB)  {
	log.Printf("Migrations started")
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("[ERROR] on db migrations (model: user)")
	}
	err = db.AutoMigrate(&models.Language{})
	if err != nil {
		log.Printf("[ERROR] on db migrations (model: language)")
	}
	err = db.AutoMigrate(&models.Word{})
	if err != nil {
		log.Printf("[ERROR] on db migrations (model: word) - %v", err.Error())
	}
	log.Printf("Migrations finished")
}

func afterConnectionWorks(db *gorm.DB)  {
	// ImplementLanguages("./internal/database/languages.json", db)
}

