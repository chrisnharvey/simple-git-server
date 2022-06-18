package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func CreateRepo() *cobra.Command {
	return &cobra.Command{
		Use:   "create-repo [name]",
		Short: "Create new git repo.",
		Run:   createRepoCommand,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("name must be provided")
			}

			return nil
		},
	}
}

func createRepoCommand(cmd *cobra.Command, args []string) {
	for _, v := range args {
		if _, err := os.Stat("/git/repos/" + v); err == nil {
			fmt.Println(errors.New(v + " already exists"))
		}
	}

	for _, v := range args {
		err := createRepo(v)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(v + " created")
	}
}

func createRepo(n string) error {
	err := runCommand("sudo", "mkdir", "/git/repos/"+n)
	if err != nil {
		return err
	}

	err = runCommand("sudo", "git", "init", "--bare", "/git/repos/"+n)
	if err != nil {
		return err
	}

	err = runCommand("sudo", "chown", "git:git", "-R", "/git/repos/"+n)
	if err != nil {
		return err
	}

	return nil
}

func runCommand(c string, args ...string) error {
	cmd := exec.Command(c, args...)

	return cmd.Run()
}
