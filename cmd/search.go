package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/stikypiston/jots/internal/storage"
	"github.com/stikypiston/jots/internal/ui"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search entries",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		query := strings.ToLower(args[0])

		entries, err := storage.ScanEntries()
		if err != nil {
			return err
		}

		for _, e := range entries {
			if strings.Contains(strings.ToLower(e.Content), query) {
				fmt.Println(ui.FormatEntry(e))
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
