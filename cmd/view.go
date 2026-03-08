package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/stikypiston/jots/internal/storage"
	"github.com/stikypiston/jots/internal/ui"
)

var viewCmd = &cobra.Command{
	Use:   "view [uuid]",
	Short: "View an entry",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		id := args[0]

		entries, err := storage.ScanEntries()
		if err != nil {
			return err
		}

		for _, e := range entries {

			if strings.HasPrefix(e.ID.String(), id) {

				fmt.Println(ui.FormatEntry(e))

				if len(e.Attachments) > 0 {
					fmt.Println()
					fmt.Println("Attachments:")

					for _, a := range e.Attachments {
						fmt.Printf("  %s (%s)\n",
							a.Filename,
							a.ID.String()[:8],
						)
					}
				}

				return nil
			}
		}

		return fmt.Errorf("entry not found")
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}
