package main

import (
	"log"

	"paddletraffic/internal/repository"
)

func main() {
	db, err := repository.OpenDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connected successfully")

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()
}
