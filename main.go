package main

import (
	"log"
)

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
