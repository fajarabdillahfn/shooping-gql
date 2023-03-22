package repository

import (
	"log"

	rShop "github.com/fajarabdillahfn/shoping-gql/internal/repository"
	"gorm.io/gorm"
)

type repository struct {
	conn *gorm.DB
}

func NewShopRepo(db *gorm.DB) rShop.Repository {
	if db == nil {
		log.Panic("missing database connection")
	}

	repo := &repository{
		conn: db,
	}

	return repo
}
