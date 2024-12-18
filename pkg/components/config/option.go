package config

import (
	"fmt"
)

// get the default options
func GetDefaultOptions() []Option {
	return []Option{
		{Label: "fix", Value: "fix", Usage: 0},
		{Label: "feat", Value: "feat", Usage: 0},
		{Label: "refactor", Value: "refactor", Usage: 0},
		{Label: "build", Value: "build", Usage: 0},
		{Label: "chore", Value: "chore", Usage: 0},
		{Label: "style", Value: "style", Usage: 0},
		{Label: "docs", Value: "docs", Usage: 0},
		{Label: "revert", Value: "revert", Usage: 0},
		{Label: "test", Value: "test", Usage: 0},
	}
}

func GetOptions() ([]Option, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT label, value, usage FROM options ORDER BY usage DESC")
	if err != nil {
		fmt.Println("❌ line 31 err ➡️", err)
		return nil, fmt.Errorf("failed to query options: %w", err)
	}
	defer rows.Close()

	var options []Option
	for rows.Next() {
		var option Option
		if err := rows.Scan(&option.Label, &option.Value, &option.Usage); err != nil {
			fmt.Println("❌ line 40 err ➡️", err)
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		options = append(options, option)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("❌ line 47 err ➡️", err)
		return nil, fmt.Errorf("error during rows iteration: %w", err)
	}

	return options, nil
}

func SaveOptions(options []Option) error {
	return SaveRecords(options, "options", []string{"label", "value", "usage"}, "value", []string{"label", "usage"})
}
