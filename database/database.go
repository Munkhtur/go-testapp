package initializers

import (
	"example/testapp/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectToDB() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection failed")
	}

	db.AutoMigrate(&models.User{}, models.Word{}, models.Synonym{})

	DB = Dbinstance{
		Db: db,
	}
}
