package database

import (
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/crypto-platform/shared/config"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CloseConnection func() error

func ConnectPostgres(config config.Postgres) (*gorm.DB, CloseConnection, error) {

	db, err := gorm.Open(postgres.Open(connString(config)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt: true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get database connection: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, nil, fmt.Errorf("failed to ping database: %w", err)
	}

	closeFunc := func() error {
		log.Println("Closing PostgreSQL database connection")
		if err := sqlDB.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %w", err)
		}
		log.Println("PostgreSQL database connection closed successfully")
		return nil
	}
	log.Println("Connected to PostgreSQL database")

	return db, closeFunc, nil
}

func connString(config config.Postgres) string {
	return fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Username, config.Password, config.Database, config.SSLMode)
}
