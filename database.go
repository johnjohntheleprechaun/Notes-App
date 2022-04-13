package main

import (
	"database/sql"
	"os"
	"github.com/mattn/go-sqlite3"
)

func initDB(filepath string) (error){
	//os.Remove(filepath)
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {return err}
	_, _ = db.Prepare(
`CREATE TABLE IF NOT EXISTS users (
	uuid INTEGER PRIMARY KEY,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL UNIQUE
)`)
	return nil
}