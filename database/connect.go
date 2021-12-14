package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/arunap2509/notes-api/config"
	"github.com/arunap2509/notes-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Print("something went wrong")
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("database connected")

	DB.AutoMigrate(&model.Note{}, &model.User{})

	// _ = DB.Migrator().DropColumn(&model.Note{}, "can_delete")

	fmt.Println("database migration completed")
}
