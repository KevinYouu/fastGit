package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

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

func SaveRecords[T any](
	records []T,
	tableName string,
	columns []string,
	conflictColumn string, // If conflictColumn is not empty, use ON CONFLICT
	updateColumns []string, // Columns that need to be updated during a conflict
) error {
	db, err := openDB()
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// 动态构建插入 SQL
	columnPlaceholders := make([]string, len(columns))
	for i := range columns {
		columnPlaceholders[i] = "?"
	}

	// 基础插入 SQL
	baseSQL := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(columnPlaceholders, ", "),
	)

	// 添加 ON CONFLICT 子句（如果需要）
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

	// 准备语句
	stmt, err := tx.Prepare(baseSQL)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// 执行插入或更新
	for _, record := range records {
		values := extractValues(record, columns)
		_, err := stmt.Exec(values...)
		if err != nil {
			return fmt.Errorf("failed to execute statement: %w", err)
		}
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// extractValues 是一个辅助函数，用于从结构体中提取列值
func extractValues[T any](record T, columns []string) []interface{} {
	v := reflect.ValueOf(record)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	values := make([]interface{}, len(columns))
	for i, col := range columns {
		field := v.FieldByName(strings.Title(col)) // 假设字段名与列名匹配
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
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	// Create the patches table
	if err := migratePatchesTable(db); err != nil {
		fmt.Println("❌ line 104 err ➡️", err)
		return fmt.Errorf("failed to migrate table: %w", err)
	}

	// Count the number of options
	countQuery := `SELECT COUNT(*) FROM options`
	var count int
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		fmt.Println("❌ line 113 err ➡️", err)
		return fmt.Errorf("failed to count options: %w", err)
	}

	// If no options exist, insert default options
	if count == 0 {
		defaultOptions := GetDefaultOptions()
		fmt.Println(defaultOptions)
		err = SaveOptions(defaultOptions)
		if err != nil {
			fmt.Println("❌ line 123 err ➡️", err)
			return fmt.Errorf("failed to insert default options: %w", err)
		}
	}

	// Count the number of patches
	countQuery = `SELECT COUNT(*) FROM patches`
	err = db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		fmt.Println("❌ line 132 err ➡️", err)
		return fmt.Errorf("failed to count patches: %w", err)
	}

	// If no patches exist, insert default patches
	if count == 0 {
		defaultPatches := GetDefaultTagPatch()
		fmt.Println("🚀 line 136 defaultPatches ➡️", defaultPatches)
		err = SavePatches(defaultPatches)
		if err != nil {
			return fmt.Errorf("failed to insert default patches: %w", err)
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
		return fmt.Errorf("failed to check table existence for %s: %w", tableName, err)
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
			return fmt.Errorf("failed to create table %s: %w", tableName, err)
		}
	} else if err == nil && tableExists == 1 {
		// If the table exists, check for missing columns
		existingColumns := make(map[string]bool)
		tableInfoQuery := fmt.Sprintf(`PRAGMA table_info(%s)`, tableName)
		rows, err := db.Query(tableInfoQuery)
		if err != nil {
			return fmt.Errorf("failed to query table info for %s: %w", tableName, err)
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
				return fmt.Errorf("failed to scan table info for %s: %w", tableName, err)
			}
			existingColumns[name] = true
		}

		// Add missing columns
		for column, definition := range requiredColumns {
			if !existingColumns[column] {
				alterQuery := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s", tableName, column, definition)
				if _, err := db.Exec(alterQuery); err != nil {
					return fmt.Errorf("failed to add column %s to table %s: %w", column, tableName, err)
				}
			}
		}
	}

	return nil
}
