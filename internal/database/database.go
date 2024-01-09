package database

import (
	"errors"
	"fmt"
	"go-simple-template/config"
	"go-simple-template/pkg/logger"

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
		log = logger.NewLogger().Logger.With().Str("pkg", "main").Logger()
	)

	gormConfig := &gorm.Config{}

	switch cfg.DBconfig.DBdriver {
	case driverMySQL:
		dsn = fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s`,
			cfg.DBconfig.DBuser,
			cfg.DBconfig.DBpassword,
			cfg.DBconfig.DBhost,
			cfg.DBconfig.DBport,
			cfg.DBconfig.DBname,
		)

		dbConn, err := gorm.Open(mysql.Open(dsn), gormConfig)
		if err != nil {
			log.Error().Err(err).Str("dsn", dsn).Msg("failed to connect to database")

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
