package main

import (
	"exercisee/db"
	"log"
)

func main() {
	err := db.InitDB()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer db.CloseConnection()
}
