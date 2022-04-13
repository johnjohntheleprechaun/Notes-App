package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const userTableCreate = `
CREATE TABLE IF NOT EXISTS Users (
	UserID INTEGER PRIMARY KEY,
	Username TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL
)`
const notesTableCreate = `
CREATE TABLE IF NOT EXISTS Notes (
	NoteID INTEGER PRIMARY KEY,
	Owner INTEGER NOT NULL,
	Content TEXT,
	FOREIGN KEY(Owner) REFERENCES Users(UserID)
)`

func initDB(filepath string) *sql.DB {
	//os.Remove(filepath)
	db, err := sql.Open("sqlite3", filepath)
	checkError(err)
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	_, err = db.Exec(userTableCreate)
	_, err = db.Exec(notesTableCreate)
	return db
}
