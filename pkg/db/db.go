package db

import (
	"database/sql"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const dbEnv = "DB_FILE"

func init() {
	dbfile := os.Getenv(dbEnv)
	if dbfile == "" {
		log.Printf("DB_FILE environment variable not set, using default db file")
	}

	pwd, _ := os.Getwd()
	// TODO: read sql schema file instead of db file
	dbfile = path.Join(pwd, "./watchdog.sqlite")
	log.Printf("Using db file %s", dbfile)
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	log.Printf("connected to db: %s", dbfile)
}
