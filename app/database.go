package app

import (
	"be_tabungan_api/entity"
	"database/sql"
	"fmt"
	"time"
)

func NewDB(dbConfig entity.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", dbConfig.DBUser, dbConfig.DBPassword, dbConfig.DBHost, dbConfig.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
