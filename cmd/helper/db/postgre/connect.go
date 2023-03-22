package postgre

import (
	"fmt"
	"log"
	"os"

	"github.com/fajarabdillahfn/shoping-gql/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB() *gorm.DB {
	log.Println("connecting to db...")

	host := os.Getenv("SQL_HOST")
	port := os.Getenv("SQL_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	openConnection := postgres.Open(psqlInfo)

	DB, err := gorm.Open(openConnection, &gorm.Config{})
	DB = DB.Set("gorm:auto_preload", true)
	if err != nil {
		panic("failed to connect database")
	}

	if err := autoMigrate(DB); err != nil {
		panic("failed to migrate, caused: " + err.Error())
	}

	log.Println("db connected")
	return DB
}

func autoMigrate(DB *gorm.DB) error {
	err := DB.AutoMigrate(&model.Product{})

	if err != nil {
		return err
	}

	return nil
}
