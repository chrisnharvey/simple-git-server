package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func CreateRepo() *cobra.Command {
	return &cobra.Command{
		Use:       "create-repo",
		Short:     "Create new git repo.",
		ValidArgs: []string{"a", "b", "c"},
		Run:       createRepoCommand,
	}
}

func createRepoCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Hello world")
}
