package config

import "fmt"

func IncrementUsage(value string) error {
	db, err := openDB()
	if err != nil {
		fmt.Println("❌ line 8 err ➡️", err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println("❌ line 15 err ➡️", err)
		return err
	}
	defer tx.Rollback() // Ensure transaction is rolled back on error

	// Prepare the statement to update the usage count
	stmt, err := tx.Prepare(`
		UPDATE options
		SET usage = usage + 1
		WHERE value = ?
	`)
	if err != nil {
		fmt.Println("❌ line 27 err ➡️", err)
		return err
	}
	defer stmt.Close()

	// Execute the statement
	_, err = stmt.Exec(value)
	if err != nil {
		tx.Rollback() // Rollback on error
		fmt.Println("❌ line 36 err ➡️", err)
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		fmt.Println("❌ line 42 err ➡️", err)
		return err
	}

	return nil
}
