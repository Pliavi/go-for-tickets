package database

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db   *sqlx.DB
	once sync.Once
)

func MustGetConnection() *sqlx.DB {
	once.Do(func() {
		pguser := os.Getenv("PG_USER")
		pgdb := os.Getenv("PG_DB")
		pghost := os.Getenv("PG_HOST")
		pgport := os.Getenv("PG_PORT")
		pgpass := os.Getenv("PG_PASS")

		dbURI := fmt.Sprintf("user=%s dbname=%s host=%s port=%v sslmode=disable", pguser, pgdb, pghost, pgport)
		if pgpass != "" {
			dbURI += " password=" + pgpass
		}
		var err error

		db, err = sqlx.Connect("postgres", dbURI)
		if err != nil {
			panic(fmt.Sprintf("Unable to connection to database: %v\n", err))
		}

		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
	})

	return db
}
