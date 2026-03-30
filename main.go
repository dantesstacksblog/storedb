package main

import (
	"log"
	"storedb/db"
)

func main() {
	database, err := db.New("store.db?_foreign_keys=on")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer database.Close()

	if err := database.Migrate(); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	log.Println("StoreDB is running — schema ready")
}
