package databases

import (
	"log"
	"tder/internal/entities"

	"gorm.io/gorm"
)

func InitDb(db *gorm.DB) {
	models := []interface{}{
		&entities.User{},
	}

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Fatalf("Error migrating database schema: %v", err)
		}
	}

}
