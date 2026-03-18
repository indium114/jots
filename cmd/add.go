package cmd

import (
	"strings"
	"time"

	"github.com/charmbracelet/huh"
	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/stikypiston/jots/internal/models"
	"github.com/stikypiston/jots/internal/storage"
)

var attachmentFlag string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a journal entry",
	RunE: func(cmd *cobra.Command, args []string) error {

		storage.EnsureDirs()

		var content string
		var attach bool
		var attachmentPath string

		// non-interactive mode
		if len(args) > 0 {
			content = strings.Join(args, " ")
		} else {
			form := huh.NewForm(
				huh.NewGroup(
					huh.NewText().
						Title("Entry text").
						Value(&content),

					huh.NewConfirm().
						Title("Add attachment?").
						Value(&attach),
				),
			)

			err := form.Run()
			if err != nil {
				return err
			}
		}

		var attachments []models.Attachment

		if attachmentFlag != "" {
			a, err := storage.CopyAttachment(attachmentFlag)
			if err != nil {
				return err
			}

			attachments = append(attachments, a)
		}

		if attach && attachmentFlag == "" {
			err := huh.NewInput().
				Title("Attachment path").
				Value(&attachmentPath).
				Run()

			if err != nil {
				return err
			}

			a, err := storage.CopyAttachment(attachmentPath)
			if err != nil {
				return err
			}

			attachments = append(attachments, a)
		}

		now := time.Now()

		entry := models.Entry{
			ID:          uuid.New(),
			Timestamp:   now,
			Content:     content,
			Attachments: attachments,
		}

		df, err := storage.LoadDay(now)
		if err != nil {
			return err
		}

		df.Entries = append(df.Entries, entry)

		return storage.SaveDay(now, df)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&attachmentFlag, "attachment", "a", "", "Absolute path to attachment")
}
