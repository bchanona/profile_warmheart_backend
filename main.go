package main

import "github.com/bchanona/profile_warmheart_backend/helpers"

func main() {
	db, _ := helpers.ConnectMySQL()
	if db == nil {
		panic("Error connecting to the database")
	} else {
		println("Successful connection to the database")
	}
}
