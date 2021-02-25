package cmd

import (
	"errors"
	"time"

	"github.com/b4b4r07/history/history"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add new history",
	Long:  "Add new history",
	RunE:  add,
}

func add(cmd *cobra.Command, args []string) error {
	h, err := history.Load()
	if err != nil {
		return err
	}

	r := history.NewRecord()
	// Skip adding if the command is registed as ignoring word
	if history.CheckIgnores(addCommand) {
		return nil
	}

	r.SetCommand(addCommand)
	r.SetDir(addDir)
	r.SetBranch(addBranch)
	r.SetStatus(addStatus)

	// Backup before adding new record
	// However don't backup many times on the same day
	if h.Records.Latest().Date.Day() != time.Now().Day() {
		if err := h.Backup(); err != nil {
			return err
		}
	}
	h.Records.Add(*r)

	return h.Save()
}

var (
	addCommand string
	addDir     string
	addBranch  string
	addStatus  int
)

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&addCommand, "command", "", "", "Set command")
	addCmd.Flags().StringVarP(&addDir, "dir", "", "", "Set dir")
	addCmd.Flags().StringVarP(&addBranch, "branch", "", "", "Set branch")
	addCmd.Flags().IntVarP(&addStatus, "status", "", 0, "Set status")
}
