package migration

import (
	"go-simple-template/internal/database"
	"log"
)

func AutoMigrate() error {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	tables := []interface{}{
		// Add your tables here
	}

	err = db.AutoMigrate(tables...)
	if err != nil {
		return err
	}

	return nil
}
