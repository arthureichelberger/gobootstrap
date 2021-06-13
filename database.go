package gobootstrap

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres support.
)

type DatabaseConnection struct {
	Db *sqlx.DB
}

func NewDatabaseConnection(host string, port string, username string, password string, dbname string) (DatabaseConnection, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return DatabaseConnection{}, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	err = db.Ping()
	if err != nil {
		return DatabaseConnection{}, err
	}

	return DatabaseConnection{
		Db: db,
	}, nil
}
