package main

import (
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

//Table Creation
const userTableCreate = `
CREATE TABLE IF NOT EXISTS Users (
	UserID INTEGER PRIMARY KEY,
	Username TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL,
	AuthToken TEXT UNIQUE
)`
const notesTableCreate = `
CREATE TABLE IF NOT EXISTS Notes (
	NoteID INTEGER PRIMARY KEY,
	Owner INTEGER NOT NULL,
	Content TEXT,
	FOREIGN KEY(Owner) REFERENCES Users(UserID)
)`

//Data Insertion/Editing/Retrieval
const userCreate = "INSERT INTO Users (Username, Password) VALUES (?, ?)"
const noteCreate = "INSERT INTO Notes (Owner) VALUES (?) WHERE (SELECT AuthToken FROM Users WHERE UserID=?)=?"
const noteEdit = "UPDATE Notes SET Content=? WHERE NoteID=? AND (SELECT AuthToken FROM Users WHERE UserID=?)=?"
const noteFetch = "SELECT Note FROM Notes WHERE NoteID=? AND (SELECT AuthToken FROM Users WHERE UserID=(SELECT Owner FROM Notes WHERE NoteID=?))=?"

//Existence/Auth Checks
const userExists = "SELECT UserID FROM Users WHERE UserID=?"
const noteExists = "SELECT NoteID FROM Notes WHERE NoteID=?"
const authTokenCheck = "SELECT AuthToken FROM Users WHERE AuthToken=?"

func initDB(filepath string) *sql.DB {
	//os.Remove(filepath)
	db, err := sql.Open("sqlite3", filepath)
	checkError(err)
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	_, err = db.Exec(userTableCreate)
	_, err = db.Exec(notesTableCreate)
	return db
}

func testDB(db *sql.DB) {
	fmt.Println(createUser(db, "johnjohntheleprechaun", "password"))
	fmt.Println(checkUserExists(db, 4))
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
	if !checkUserExists(db, owner) {
		return false
	}
	result, _ := db.Exec(noteCreate, owner, owner, authToken)
	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	} else {
		return true
	}
}

func editNote(db *sql.DB, newContent string, noteID int, authToken int, owner int) bool {
	result, _ := db.Exec(noteEdit, newContent, noteID, owner, authToken)
	if affected, _ := result.RowsAffected(); affected == 0 {
		return false
	} else {
		return true
	}
}

func getNote(db *sql.DB, noteID int, authToken string, owner int) (string, bool) {
	rows, _ := db.Query(noteFetch, noteID, noteID, authToken)
	defer rows.Close()
	var note string
	if rows.Next() {
		rows.Scan(&note)
		return note, true
	} else {
		return note, false
	}
}

func checkUserExists(db *sql.DB, userID int) bool {
	rows, _ := db.Query(userExists, userID)
	if rows.Next() {
		return true
	} else {
		return false
	}
}

func checkNoteExists(db *sql.DB, noteID int) bool {
	rows, _ := db.Query(noteExists, noteID)
	if rows.Next() {
		return true
	} else {
		return false
	}
}

func validateAuthToken(db *sql.DB, authToken string) bool {
	rows, _ := db.Query(authTokenCheck, authToken)
	if rows.Next() {
		return true
	} else {
		return false
	}
}
