package main

import (
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

const userTableCreate = `
CREATE TABLE IF NOT EXISTS Users (
	UserID INTEGER PRIMARY KEY,
	Username TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL,
	AuthToken INTEGER UNIQUE
)`
const notesTableCreate = `
CREATE TABLE IF NOT EXISTS Notes (
	NoteID INTEGER PRIMARY KEY,
	Owner INTEGER NOT NULL,
	Content TEXT,
	FOREIGN KEY(Owner) REFERENCES Users(UserID)
)`
const userCreate = "INSERT INTO Users (Username, Password) VALUES (?, ?)"
const noteCreate = "INSERT INTO Notes (Owner) VALUES (?) WHERE (SELECT AuthToken FROM Users WHERE UserID=?)=?"
const noteEdit = "UPDATE Notes SET Content=? WHERE NoteID=? AND (SELECT AuthToken FROM Users WHERE UserID=?)=?"

func initDB(filepath string) *sql.DB {
	//os.Remove(filepath)
	db, err := sql.Open("sqlite3", filepath)
	checkError(err)
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	_, err = db.Exec(userTableCreate)
	_, err = db.Exec(notesTableCreate)
	fmt.Println(createUser(db, "Johnny_Thunder", "thisisapassword"))
	return db
}

func createUser(db *sql.DB, username string, password string) bool {
	_, err := db.Exec(userCreate, username, password)
	goodUsername := true
	if e, ok := err.(sqlite3.Error); ok {
		goodUsername = (e.ExtendedCode != sqlite3.ErrConstraintUnique)
	}
	return goodUsername
}

func createNote(db *sql.DB, owner int, authToken int) bool {
	result, _ := db.Exec(noteCreate, owner, owner, authToken)
	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	} else {
		return true
	}
}

func editNote(db *sql.DB, newContent string, noteID int, authToken int, owner int) {
	_, _ = db.Exec(noteEdit, newContent, noteID, owner, authToken)
}
