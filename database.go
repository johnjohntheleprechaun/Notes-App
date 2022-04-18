package main

import (
	"database/sql"

	"math/rand"

	_ "github.com/mattn/go-sqlite3"
)

const b64Chars string = "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890+="

//Table Creation
const userTableCreate string = `
CREATE TABLE IF NOT EXISTS Users (
	UserID INTEGER PRIMARY KEY,
	Username TEXT NOT NULL UNIQUE,
	Password TEXT NOT NULL,
	AuthToken TEXT UNIQUE
)`
const notesTableCreate string = `
CREATE TABLE IF NOT EXISTS Notes (
	NoteID INTEGER PRIMARY KEY,
	Owner INTEGER NOT NULL,
	Content TEXT,
	FOREIGN KEY(Owner) REFERENCES Users(UserID)
)`

//Data Insertion/Editing/Retrieval
const userCreate string = "INSERT INTO Users (Username, Password) VALUES (?, ?)"
const noteCreate string = "INSERT INTO Notes (Owner) VALUES (?) WHERE (SELECT AuthToken FROM Users WHERE UserID=?)=?"
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
}

//Authorization Management
func createUser(db *sql.DB, username string, password string) {
	db.Exec(userCreate, username, password)
}

func createAuthToken() string {
	var token string
	for i := 0; i < 10; i++ {
		token += string(b64Chars[rand.Intn(63)])
	}
	return token
}
func setNewAuthToken(db *sql.DB, userID int) {

}
