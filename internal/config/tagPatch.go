package config

import (
	"database/sql"
)

func GetDefaultTagPatch() []Patch {
	return []Patch{
		{Prefix: "", Major: 999, Minor: 9, Patch: 9, Suffix: ""},
	}
}

func GetTagPatch() (Patch, error) {
	db, err := openDB()
	if err != nil {
		return GetDefaultTagPatch()[0], err
	}
	defer db.Close()

	row := db.QueryRow("SELECT prefix, major, minor, patch, suffix FROM patches LIMIT 1")
	var patch Patch
	if err := row.Scan(&patch.Prefix, &patch.Major, &patch.Minor, &patch.Patch, &patch.Suffix); err != nil {
		if initErr := Initialize(); initErr != nil {
			return GetDefaultTagPatch()[0], initErr
		}
		row = db.QueryRow("SELECT prefix, major, minor, patch, suffix FROM patches LIMIT 1")
		if err := row.Scan(&patch.Prefix, &patch.Major, &patch.Minor, &patch.Patch, &patch.Suffix); err != nil {
			if err == sql.ErrNoRows {
				return GetDefaultTagPatch()[0], err
			}
			return GetDefaultTagPatch()[0], err
		}
	}

	return patch, nil
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
