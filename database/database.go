package database

import (
	"errors"
	"fmt"
	"go-simple-template/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	driverMySQL      = "mysql"
	driverPostgreSQL = "postgres"
)

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	var (
		dsn string
	)

	gormConfig := &gorm.Config{}

	switch cfg.DBconfig.DBdriver {
	case driverMySQL:
		dsn = fmt.Sprintf(`%s@%stcp(%s:%s)/%s`,
			cfg.DBconfig.DBuser,
			cfg.DBconfig.DBpassword,
			cfg.DBconfig.DBhost,
			cfg.DBconfig.DBport,
			cfg.DBconfig.DBname,
		)

		dbConn, err := gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			return nil, err
		}

		return dbConn, nil

	case driverPostgreSQL:
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBconfig.DBhost,
			cfg.DBconfig.DBuser,
			cfg.DBconfig.DBpassword,
			cfg.DBconfig.DBname,
			cfg.DBconfig.DBport,
		)

		dbConn, err := gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			return nil, err
		}

		return dbConn, nil
	}

	return nil, errors.New("invalid database driver")
}
