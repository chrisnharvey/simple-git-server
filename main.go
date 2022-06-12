package main

import (
	"fmt"

	shell "github.com/brianstrauch/cobra-shell"

	"github.com/chrisnharvey/simple-git-server/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cobra := &cobra.Command{}

	cobra.AddCommand(cmd.CreateRepo())

	shell := shell.New(cobra)

	err := shell.Execute()

	if err != nil {
		fmt.Println(err)
	}
}
