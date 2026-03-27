package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/indium114/jots/internal/storage"
)

var openCmd = &cobra.Command{
	Use:   "open [attachment-uuid]",
	Short: "Open an attachment by its UUID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		entries, err := storage.ScanEntries()
		if err != nil {
			return err
		}

		var attachmentPath string
		found := false

		for _, e := range entries {
			for _, a := range e.Attachments {
				if strings.HasPrefix(a.ID.String(), id) {
					attachmentPath = storage.AttachmentsDir() + "/" + a.Stored
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			return fmt.Errorf("attachment not found")
		}

		var cmdExec *exec.Cmd

		switch runtime.GOOS {
		case "darwin":
			cmdExec = exec.Command("open", attachmentPath)
		case "linux":
			cmdExec = exec.Command("xdg-open", attachmentPath)
		case "windows":
			cmdExec = exec.Command("cmd", "/c", "start", "", attachmentPath)
		default:
			return fmt.Errorf("unsupported platform")
		}

		cmdExec.Stdout = os.Stdout
		cmdExec.Stderr = os.Stderr

		return cmdExec.Run()
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
