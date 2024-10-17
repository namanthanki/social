package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/namanthanki/social/internal/db"
	"github.com/namanthanki/social/internal/env"
	"github.com/namanthanki/social/internal/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %e", err)
	}

	addr := env.GetString("DB_ADDRESS", "postgresql://postgres:password@localhost:5432/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer conn.Close()

	store := store.NewPostgresStorage(conn)

	db.Seed(store, conn)
}
