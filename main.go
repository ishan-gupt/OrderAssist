package main

import (
	"BeBot/whatsapp"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := whatsapp.Connect()
	if err != nil {
		panic(err)
	}
	defer whatsapp.Disconnect()
}
