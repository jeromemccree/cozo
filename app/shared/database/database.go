package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var BACKENDDB *sql.DB

func postgresqldsn() string {

	return "user=" + os.Getenv("PGUsername") + " " +
		"password=" + os.Getenv("PGPassword") + " " +
		"dbname=" + os.Getenv("PGName") + " " +
		"host=" + os.Getenv("PGHost") + " " +
		"port=" + os.Getenv("PGPort") + " " +
		"sslmode=" + os.Getenv("PGSSLMode")
}

func ConnectPostgresSQL() {
	var err error

	//connect to postgresql
	if BACKENDDB, err = sql.Open("postgres", postgresqldsn()); err != nil {
		log.Println("Postgres SQL Driver Error")
	}

	if err = BACKENDDB.Ping(); err != nil {
		log.Println("Database Error")
	}
}
