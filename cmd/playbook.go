/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/sebastianwebber/a/pkg/task"
	"github.com/spf13/cobra"
)

// playbookCmd represents the playbook command
var (
	playbookCmd = &cobra.Command{
		Use:   "playbook [playbook_file]",
		Args:  cobra.MinimumNArgs(1),
		Short: "Executes a playbook file that is compatible with the ansible playbook format",
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]

			p, err := task.ParsePlaybook(filename)
			if err != nil {
				logger.Fatal("Error while parsing the playbook", "error", err)
			}

			logger.Debug("Playbook parsed", "playbook", spew.Sdump(p))

		},
	}
)

func init() {
	rootCmd.AddCommand(playbookCmd)
}
