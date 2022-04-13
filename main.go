package main

import (
	"log"
)

type SaveData struct {
	Session int
	ID      int
	Content string
}

func main() {
	router := initializeRouter()
	_ = initDB("./test.db")
	router.Run()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
