package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/stikypiston/jots/internal/storage"
	"github.com/stikypiston/jots/internal/ui"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List today's entries",
	RunE: func(cmd *cobra.Command, args []string) error {

		now := time.Now()

		df, err := storage.LoadDay(now)
		if err != nil {
			return err
		}

		for _, e := range df.Entries {
			fmt.Println(ui.FormatEntry(e))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
