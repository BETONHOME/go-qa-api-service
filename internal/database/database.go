package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=postgres user=postgres password=root dbname=app port=5432 sslmode=disable"

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("БД запускается", i+1)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	if err != nil {
		log.Fatal("Не удалось подключиться к БД", err)
	}

	DB = db
	log.Println("БД подключена")
}
