package database

import (
	"errors"
	"fmt"
	"go-simple-template/config"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	driverMySQL      = "mysql"
	driverPostgreSQL = "postgres"
)

func NewConnection() (*gorm.DB, error) {
	var (
		dsn string
	)

	gormConfig := &gorm.Config{}

	switch config.DBDriver() {
	case driverMySQL:
		dsn = fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`,
			config.DBUser(),
			config.DBPassword(),
			config.DBHost(),
			config.DBPort(),
			config.DBName(),
		)

		dbConn, err := gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			log.Error().Err(err).Str("dsn", dsn).Msg("failed to connect to database")

			return nil, err
		}

		return dbConn, nil

	case driverPostgreSQL:
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			config.DBHost(),
			config.DBUser(),
			config.DBPassword(),
			config.DBName(),
			config.DBPort(),
		)

		dbConn, err := gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			log.Error().Err(err).Str("dsn", dsn).Msg("failed to connect to database")

			return nil, err
		}

		return dbConn, nil
	}

	return nil, errors.New("invalid database driver")
}

// auto migrate
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
	//  add your model here, ex : &model.User{}
	)
}
