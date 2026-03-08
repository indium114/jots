package models

import (
	"time"

	"github.com/google/uuid"
)

type Attachment struct {
	ID       uuid.UUID `json:"id"`
	Filename string    `json:"filename"`
	Stored   string    `json:"stored"`
	Mime     string    `json:"mime"`
}

type Entry struct {
	ID          uuid.UUID    `json:"id"`
	Timestamp   time.Time    `json:"timestamp"`
	Content     string       `json:"content"`
	Attachments []Attachment `json:"attachments"`
}

type DayFile struct {
	Date    string  `json:"date"`
	Entries []Entry `json:"entries"`
}
