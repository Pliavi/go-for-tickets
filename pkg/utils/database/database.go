package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pliavi/go-for-tickets/pkg/config"
)

type DatabaseSingleton interface {
	Connect() (*sql.DB, error)
}

var instance *sql.DB

func GetInstance() *sql.DB {
	if instance == nil {
		instance = connect(config.DatabaseConfig)
	}
	return instance
}

func connect(env *config.DatabaseEnv) *sql.DB {
	// TODO: maybe there is a join function for this?
	connStr := "host=" + env.Host +
		" port=" + env.Port +
		" user=" + env.User +
		" password=" + env.Pass +
		" dbname=" + env.Dbname +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	return db
}
