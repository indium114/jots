package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/indium114/jots/internal/models"
)

func dayFilePath(t time.Time) string {
	return filepath.Join(
		EntriesDir(),
		t.Format("2006"),
		t.Format("01"),
		t.Format("02")+".json",
	)
}

func LoadDay(t time.Time) (*models.DayFile, error) {
	path := dayFilePath(t)

	data, err := os.ReadFile(path)
	if err != nil {
		return &models.DayFile{
			Date:    t.Format("2006-01-02"),
			Entries: []models.Entry{},
		}, nil
	}

	var df models.DayFile
	err = json.Unmarshal(data, &df)
	if err != nil {
		return nil, err
	}

	return &df, nil
}

func SaveDay(t time.Time, df *models.DayFile) error {
	path := dayFilePath(t)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(df, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
