package config

import "fmt"

func GetDefaultTagPatch() []Patch {
	return []Patch{
		{Prefix: "", Major: 0, Minor: 0, Patch: 0, Suffix: ""},
	}
}

func GetTagPatch() ([]Patch, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT prefix, major, minor, patch, suffix FROM options ORDER BY usage DESC")
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to query options: %w", err)
	}
	defer rows.Close()

	var Patches []Patch
	for rows.Next() {
		var patch Patch
		if err := rows.Scan(&patch.Prefix, &patch.Major, &patch.Minor, &patch.Patch, &patch.Suffix); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		Patches = append(Patches, patch)
	}

	return Patches, nil
}

func SavePatches(patches []Patch) error {
	return SaveRecords(
		patches,
		"patches",
		[]string{"prefix", "major", "minor", "patch", "suffix"},
		"",
		nil, // does not need to be updated
	)
}
