package main

import (
	"fmt"

	shell "github.com/brianstrauch/cobra-shell"

	"github.com/chrisnharvey/simple-git-server/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			shell := shell.New(cmd)

			err := shell.Execute()

			if err != nil {
				fmt.Println(err)
			}
		},
	}

	cobra.AddCommand(cmd.CreateRepo())
	cobra.AddCommand(cmd.AuthorizedKeys())

	cobra.Execute()
}
