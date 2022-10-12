package connections

import (
	"log"

	"github.com/dimasokta14/gin-rest/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln((err))
	}

	db.AutoMigrate(&entities.Product{})

	return db
}
