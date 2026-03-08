package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/stikypiston/jots/internal/models"
)

func ScanEntries() ([]models.Entry, error) {

	var entries []models.Entry

	root := EntriesDir()

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) != ".json" {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}

		var df models.DayFile
		if err := json.Unmarshal(data, &df); err != nil {
			return nil
		}

		entries = append(entries, df.Entries...)

		return nil
	})

	return entries, err
}
