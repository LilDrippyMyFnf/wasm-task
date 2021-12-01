package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigRepository struct {
	DB *sql.DB
}

var instance *ConfigRepository

func newConfigRepository() *ConfigRepository {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PWD")+"@tcp("+os.Getenv("DB_ADDR")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return &ConfigRepository{
		DB: db,
	}
}

func GetConfigRepository() *ConfigRepository {
	if instance == nil {
		instance = newConfigRepository()
		fmt.Println("New config")
	}

	fmt.Println("conected")
	return instance
}
