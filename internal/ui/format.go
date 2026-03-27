package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/indium114/jots/internal/models"
)

var (
	gray = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	blue = lipgloss.NewStyle().Foreground(lipgloss.Color("33"))
)

func FormatEntry(e models.Entry) string {
	date := e.Timestamp.Format("02/01/06")
	timeStr := e.Timestamp.Format(time.Kitchen)

	shortID := strings.Split(e.ID.String(), "-")[0]

	meta := fmt.Sprintf("%s %s %s", date, timeStr, shortID)
	meta = gray.Render(meta)

	icon := ""
	if len(e.Attachments) > 0 {
		icon = blue.Render(" ") // nerd font paperclip alternative
	}

	return fmt.Sprintf("%s %s%s",
		meta,
		icon,
		e.Content,
	)
}
