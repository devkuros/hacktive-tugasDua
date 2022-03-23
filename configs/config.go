package configs

import (
	"tugasdua/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=tugasduago port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(models.Order{}, models.Item{})
	// db.AutoMigrate(models.Item{}, models.Order{})
	// db.Model(&models.Order{}).AddForeignKey("item_id", "items(item_id", "RESTRICT", "RESTRICT")

	return db
}
