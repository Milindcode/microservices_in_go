package database

import (
	"fmt"
	"log"
	"os"
	// "time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct{ 
	DB *gorm.DB
}

var DB_OBJ *Database

func InitDB() error {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	var err error
	var database Database
	psqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	
	log.Println(psqlinfo)
	database.DB, err = gorm.Open(postgres.Open(psqlinfo), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Database Connected Successfully")


	////  TO BE CHANGED
	err = database.DB.AutoMigrate(&User{})
	if err!= nil {
		log.Println("Error creating relations in Database: ", err)
	}
	
	DB_OBJ = &database
	return  nil
}