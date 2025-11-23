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

	// Ждем пока БД запустится (10 попыток по 3 секунды)
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Попытка подключения %d/10: БД еще не готова...", i+1)
			time.Sleep(3 * time.Second)
			continue
		}
		break
	}

	if err != nil {
		log.Fatal("Не удалось подключиться к БД после 10 попыток:", err)
	}

	DB = db
	log.Println("БД подключена")
}
