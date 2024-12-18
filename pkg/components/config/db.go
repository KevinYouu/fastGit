package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// Get the path to the SQLite database
func getDBPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".fastGit.db")
}

// Open the SQLite database
func openDB() (*sql.DB, error) {
	dbPath := getDBPath()
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return db, nil
}

func SaveOptions(options []Option) error {
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback() // Ensure transaction is rolled back on error

	// Prepare the statement with ON CONFLICT clause to update
	stmt, err := tx.Prepare(`
		INSERT INTO options (label, value, usage)
		VALUES (?, ?, ?)
		ON CONFLICT(value) DO UPDATE SET
			label = excluded.label,
			usage = excluded.usage
	`)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, option := range options {
		_, err := stmt.Exec(option.Label, option.Value, option.Usage)
		if err != nil {
			tx.Rollback() // Rollback on error
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func Initialize() error {
	db, err := openDB()
	if err != nil {
		fmt.Println("open db", err)
		return err
	}
	defer db.Close()

	// Create the options table
	if err := migrateTable(db); err != nil {
		fmt.Println("migrateTable", err)
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	// Count the number of options
	countQuery := `SELECT COUNT(*) FROM options`
	var count int
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to count options: %w", err)
	}
	fmt.Println("count", count)
	// If no options exist, insert default options
	if count == 0 {
		defaultOptions := GetDefaultOptions()
		fmt.Println(defaultOptions)
		err = SaveOptions(defaultOptions)
		if err != nil {
			return fmt.Errorf("failed to insert default options: %w", err)
		}
	}

	return nil
}

func migrateTable(db *sql.DB) error {
	// Check if the table exists
	var tableExists int
	err := db.QueryRow(`SELECT 1 FROM sqlite_master WHERE type='table' AND name='options'`).Scan(&tableExists)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("failed to check table existence: %w", err)
	}

	// If table does not exist, create it
	if err == sql.ErrNoRows {
		createTableQuery := `
		CREATE TABLE options (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			label TEXT NOT NULL,
			value TEXT NOT NULL UNIQUE,
			usage INTEGER DEFAULT 0
		)`
		_, err := db.Exec(createTableQuery)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
	} else if err == nil && tableExists == 1 {
		// Table exists but does not have the correct schema
		existingColumns := make(map[string]bool)
		query := `PRAGMA table_info(options)`
		rows, err := db.Query(query)
		if err != nil {
			return fmt.Errorf("failed to query table info: %w", err)
		}
		defer rows.Close()

		// Store existing columns
		for rows.Next() {
			var (
				cid          int
				name         string
				dataType     string
				notNull      int
				defaultValue sql.NullString
				primaryKey   int
			)
			if err := rows.Scan(&cid, &name, &dataType, &notNull, &defaultValue, &primaryKey); err != nil {
				return fmt.Errorf("failed to scan table info: %w", err)
			}
			existingColumns[name] = true
		}

		// Define required columns
		requiredColumns := map[string]string{
			"id":    "INTEGER PRIMARY KEY AUTOINCREMENT",
			"label": "TEXT NOT NULL",
			"value": "TEXT NOT NULL UNIQUE",
			"usage": "INTEGER DEFAULT 0",
		}

		// Check for missing columns and add them
		for column, definition := range requiredColumns {
			if !existingColumns[column] {
				alterQuery := fmt.Sprintf("ALTER TABLE options ADD COLUMN %s %s", column, definition)
				if _, err := db.Exec(alterQuery); err != nil {
					return fmt.Errorf("failed to add column %s: %w", column, err)
				}
			}
		}
	}

	return nil
}
