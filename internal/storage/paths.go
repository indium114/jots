package storage

import (
	"os"
	"path/filepath"
)

func BaseDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".jots")
}

func EntriesDir() string {
	return filepath.Join(BaseDir(), "entries")
}

func AttachmentsDir() string {
	return filepath.Join(BaseDir(), "attachments")
}

func EnsureDirs() error {
	dirs := []string{
		EntriesDir(),
		AttachmentsDir(),
	}

	for _, d := range dirs {
		err := os.MkdirAll(d, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
