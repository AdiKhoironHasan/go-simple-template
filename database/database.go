package database

import (
	"errors"
	"fmt"
	"go-simple/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	driverMySQL = "mysql"
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
	}

	return nil, errors.New("invalid database driver")
}
