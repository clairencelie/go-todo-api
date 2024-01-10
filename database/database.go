package database

import (
	"database/sql"
	"time"

	"github.com/spf13/viper"
)

func NewDB(configPath string, isTest bool) (*sql.DB, error) {
	config := viper.New()
	config.SetConfigName("config")
	config.AddConfigPath(configPath)

	err := config.ReadInConfig()

	if err != nil {
		return nil, err
	}

	dbHost := config.GetString("DATABASE_HOST")
	dbPort := config.GetString("DATABASE_PORT")
	dbName := config.GetString("DATABASE_NAME")

	if isTest {
		dbName = config.GetString("DATABASE_TEST")
	}

	dbUser := config.GetString("DATABASE_USER")
	dbPassword := config.GetString("DATABASE_PASSWORD")
	dbProtocol := config.GetString("DATABASE_PROTOCOL")

	dsn := ""

	if dbPassword == "" {
		dsn = dbUser + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")" + "/" + dbName
	} else {
		dsn = dbUser + ":" + dbPassword + "@" + dbProtocol + "(" + dbHost + ":" + dbPort + ")" + "/" + dbName
	}

	db, errSql := sql.Open("mysql", dsn)

	if errSql != nil {
		return nil, errSql
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(15 * time.Minute)

	return db, nil
}
