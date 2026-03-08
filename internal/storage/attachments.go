package storage

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/stikypiston/jots/internal/models"
)

func CopyAttachment(path string) (models.Attachment, error) {
	id := uuid.New()

	file, err := os.Open(path)
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
