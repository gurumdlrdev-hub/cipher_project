package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func initDBs() {
	var err error

	// 1. App_user Database connection
	userConntStr := "postgres://postgres:240907@localhost:5433/app_user?sslmode=disable"
	DBs.UserDB, err = sql.Open("postgres", userConntStr)
	if err != nil {
		log.Fatalf("Error opening users DB: %v", err)
	}
	if err = DBs.UserDB.Ping(); err != nil {
		log.Fatalf("Failed to ping users DB: %v", err)
	}
	fmt.Println("Connected to app_user database!")

	// 2. Cipher Project Database connection (Where cipher_project table lives)
	cipherConntStr := "postgres://postgres:240907@localhost:5433/cipher_project?sslmode=disable"
	DBs.CipherDB, err = sql.Open("postgres", cipherConntStr)
	if err != nil {
		log.Fatalf("Error opening cipher DB: %v", err)
	}
	if err = DBs.CipherDB.Ping(); err != nil {
		log.Fatalf("Failed to ping cipher DB: %v", err)
	}
	fmt.Println("Connected to cipher_project database!")
}
