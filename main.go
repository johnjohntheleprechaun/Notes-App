package main

import (
	"log"
)

func main() {
	router := initializeRouter()
	db := initDB("./test.db")
	testDB(db)
	router.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
