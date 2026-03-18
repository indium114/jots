package storage

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/stikypiston/jots/internal/models"
)

func canonicalPath(p string) (string, error) {
	if p == "" {
		return "", nil
	}

	// Expand ~
	if strings.HasPrefix(p, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		p = filepath.Join(home, strings.TrimPrefix(p, "~"))
	}

	// Make absolute
	abs, err := filepath.Abs(p)
	if err != nil {
		return "", err
	}

	// Resolve symlinks (optional but recommended)
	resolved, err := filepath.EvalSymlinks(abs)
	if err == nil {
		abs = resolved
	}

	// Clean path (remove ./, ../, etc.)
	return filepath.Clean(abs), nil
}

func CopyAttachment(path string) (models.Attachment, error) {
	id := uuid.New()

	newPath, err := canonicalPath(path)
	if err != nil {
		return models.Attachment{}, err
	}

	file, err := os.Open(newPath)
	if err != nil {
		return models.Attachment{}, err
	}
	defer file.Close()

	buf := make([]byte, 512)
	file.Read(buf)

	mime := http.DetectContentType(buf)

	ext := filepath.Ext(path)
	stored := id.String() + ext

	destPath := filepath.Join(AttachmentsDir(), stored)

	file.Seek(0, 0)

	dest, err := os.Create(destPath)
	if err != nil {
		return models.Attachment{}, err
	}
	defer dest.Close()

	io.Copy(dest, file)

	return models.Attachment{
		ID:       id,
		Filename: filepath.Base(path),
		Stored:   stored,
		Mime:     mime,
	}, nil
}
