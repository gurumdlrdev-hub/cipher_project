package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Mapping entry matching the JSON key names
type Encode struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
	Code   string `json:"code"`
}

// Wrapper matching the JSON key "mappings"
type MatchJSON struct {
	Mappings []Encode `json:"mappings"`
}

func seedDatabase() {
	// Connection string
	connStr := "postgres://postgres:240907@localhost:5433/cipher_project?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected!")

	// Read JSON file (os.ReadFile replaces deprecated ioutil.ReadFile)
	file, err := os.ReadFile("database/cipher_mapping_1.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully read JSON file")

	var matchdata MatchJSON
	if err = json.Unmarshal(file, &matchdata); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d records to insert.\n", len(matchdata.Mappings))

	// Insert query target set to cipher_project table
	query := `INSERT INTO cipher_project ("serialnumber", "decodenumber", "code") VALUES ($1, $2, $3)`

	// Use a transaction for fast batch insertion
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	for _, item := range matchdata.Mappings {
		_, err = stmt.Exec(item.ID, item.Number, item.Code)
		if err != nil {
			tx.Rollback()
			log.Fatalf("Failed to insert record ID %d: %v", item.ID, err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All records inserted successfully into cipher_project!")
}
