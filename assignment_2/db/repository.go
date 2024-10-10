package db

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		age INTEGER NOT NULL CHECK (age > 0)
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully: users")
}

func InsertUser(db *sql.DB, name string, age int) {
	query := `INSERT INTO users (name, age) VALUES ($1, $2)`

	_, err := db.Exec(query, name, age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User added successfully: %s", name)
	fmt.Println()

}

func QueryUsers(db *sql.DB) error {
	query := "SELECT id, name, age FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	fmt.Println("users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			return fmt.Errorf("error scanning user: %v", err)
		}
		fmt.Printf("id: %d \nName: %s \nAge: %d \n", id, name, age)
	}
	return nil
}

func DeleteTable(db *sql.DB) error {
	query := `DROP TABLE users;`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error deleting table: %v", err)
	}

	fmt.Println("Table 'users' dropped successfully.")
	return nil
}
