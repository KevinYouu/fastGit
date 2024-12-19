package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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
		return nil, err
	}
	return db, nil
}

func SaveRecords[T any](
	records []T,
	tableName string,
	columns []string,
	conflictColumn string, // If conflictColumn is not empty, use ON CONFLICT
	updateColumns []string, // Columns that need to be updated during a conflict
) error {
	db, err := openDB()
	if err != nil {
		fmt.Println("❌ 41 err ➡️", err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("❌ 48 err ➡️", err)
		return err
	}
	defer tx.Rollback()

	columnPlaceholders := make([]string, len(columns))
	for i := range columns {
		columnPlaceholders[i] = "?"
	}

	baseSQL := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(columnPlaceholders, ", "),
	)

	if conflictColumn != "" {
		setClauses := make([]string, len(updateColumns))
		for i, col := range updateColumns {
			setClauses[i] = fmt.Sprintf("%s = excluded.%s", col, col)
		}
		baseSQL += fmt.Sprintf(
			" ON CONFLICT(%s) DO UPDATE SET %s",
			conflictColumn,
			strings.Join(setClauses, ", "),
		)
	}

	stmt, err := tx.Prepare(baseSQL)
	if err != nil {
		fmt.Println("❌ line 83 err ➡️", err)
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		values := extractValues(record, columns)
		_, err := stmt.Exec(values...)
		if err != nil {
			fmt.Println("❌ line 88 err ➡️", err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("❌ line 94 err ➡️", err)
		return err
	}

	return nil
}

// extractValues is a helper function to extract values from a struct
func extractValues[T any](record T, columns []string) []interface{} {
	v := reflect.ValueOf(record)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	caser := cases.Title(language.Und)

	values := make([]interface{}, len(columns))
	for i, col := range columns {
		field := v.FieldByName(caser.String(col))
		if field.IsValid() {
			values[i] = field.Interface()
		} else {
			values[i] = nil
		}
	}
	return values
}

func Initialize() error {
	db, err := openDB()
	if err != nil {
		fmt.Println("❌ line 91 err ➡️", err)
		return err
	}
	defer db.Close()

	// Create the options table
	if err := migrateOptionsTable(db); err != nil {
		fmt.Println("❌ line 98 err ➡️", err)
		return err
	}

	// Create the patches table
	if err := migratePatchesTable(db); err != nil {
		fmt.Println("❌ line 104 err ➡️", err)
		return err
	}

	// Count the number of options
	countQuery := `SELECT COUNT(*) FROM options`
	var count int
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		fmt.Println("❌ line 113 err ➡️", err)
		return err
	}

	// If no options exist, insert default options
	if count == 0 {
		defaultOptions := GetDefaultOptions()
		fmt.Println(defaultOptions)
		err = SaveOptions(defaultOptions)
		if err != nil {
			fmt.Println("❌ line 123 err ➡️", err)
			return err
		}
	}

	// Count the number of patches
	countQuery = `SELECT COUNT(*) FROM patches`
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		fmt.Println("❌ line 132 err ➡️", err)
		return err
	}

	// If no patches exist, insert default patches
	if count == 0 {
		defaultPatches := GetDefaultTagPatch()
		err = SavePatches(defaultPatches)
		if err != nil {
			fmt.Println("❌ line 179 err ➡️", err)
			return err
		}
	}

	return nil
}

func migrateOptionsTable(db *sql.DB) error {
	requiredColumns := map[string]string{
		"id":    "INTEGER PRIMARY KEY AUTOINCREMENT",
		"label": "TEXT NOT NULL",
		"value": "TEXT NOT NULL UNIQUE",
		"usage": "INTEGER DEFAULT 0",
	}
	return migrateTable(db, "options", requiredColumns)
}

func migratePatchesTable(db *sql.DB) error {
	requiredColumns := map[string]string{
		"id":     "INTEGER PRIMARY KEY AUTOINCREMENT",
		"prefix": "TEXT",
		"major":  "INTEGER NOT NULL",
		"minor":  "INTEGER NOT NULL",
		"patch":  "INTEGER NOT NULL",
		"suffix": "TEXT",
	}
	return migrateTable(db, "patches", requiredColumns)
}

func migrateTable(db *sql.DB, tableName string, requiredColumns map[string]string) error {
	// Check if the table exists
	var tableExists int
	query := fmt.Sprintf(`SELECT 1 FROM sqlite_master WHERE type='table' AND name='%s'`, tableName)
	err := db.QueryRow(query).Scan(&tableExists)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("❌ line 215 err ➡️", err)
		return err
	}

	// If the table doesn't exist, create it
	if err == sql.ErrNoRows {
		var columnDefinitions []string
		for column, definition := range requiredColumns {
			columnDefinitions = append(columnDefinitions, fmt.Sprintf("%s %s", column, definition))
		}
		createTableQuery := fmt.Sprintf(
			"CREATE TABLE %s (%s)",
			tableName,
			strings.Join(columnDefinitions, ", "),
		)
		_, err := db.Exec(createTableQuery)
		if err != nil {
			fmt.Println("❌ line 232 err ➡️", err)
			return err
		}
	} else if err == nil && tableExists == 1 {
		// If the table exists, check for missing columns
		existingColumns := make(map[string]bool)
		tableInfoQuery := fmt.Sprintf(`PRAGMA table_info(%s)`, tableName)
		rows, err := db.Query(tableInfoQuery)
		if err != nil {
			fmt.Println("❌ line 241 err ➡️", err)
			return err
		}
		defer rows.Close()

		// Build a map of existing columns
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
				fmt.Println("❌ line 257 err ➡️", err)
				return err
			}
			existingColumns[name] = true
		}

		// Add missing columns
		for column, definition := range requiredColumns {
			if !existingColumns[column] {
				alterQuery := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tableName, column, definition)
				if _, err := db.Exec(alterQuery); err != nil {
					fmt.Println("❌ line 268 err ➡️", err)
					return err
				}
			}
		}
	}

	return nil
}
