package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
	"os"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("db"))
	logFatal(err)

	db, err := sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
