package db

import (
    "database/sql"
    "fmt"
    "log"
 _  "github.com/mattn/go-sqlite3"

)

type DB struct {
    *sql.DB
}

func New(path string) (*DB, error) {
    sqldb, err := sql.Open("sqlite3", path)
    if err != nil {
        return nil, fmt.Errorf("opening db: %w", err)
    }

    if err := sqldb.Ping(); err != nil {
        return nil, fmt.Errorf("connecting to db: %w", err)
    }

    log.Printf("connected to database: %s", path)
    return &DB{sqldb}, nil
}

func (d *DB) Close() error {
    return d.DB.Close()
}