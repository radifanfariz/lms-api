package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func ConnectToDB() {
	var err error

	host := os.Getenv("DB_Host")
	user := os.Getenv("DB_User")
	password := os.Getenv("DB_Password")
	dbName := os.Getenv("DB_DBName")
	port := os.Getenv("DB_Port")

	dbConfig := DBConfig{
		Host:     host,
		User:     user,
		Password: password,
		DBName:   dbName,
		Port:     port,
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
}
