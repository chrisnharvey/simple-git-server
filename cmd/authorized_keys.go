package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

func AuthorizedKeys() *cobra.Command {
	return &cobra.Command{
		Use:   "authorized-keys [user]",
		Short: "List the authorized keys for the given user.",
		Run:   listAuthorizedKeys,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("user must be provided")
			}

			if len(args) > 1 {
				return errors.New("too many arguments provided")
			}

			return nil
		},
	}
}

func listAuthorizedKeys(cmd *cobra.Command, args []string) {
	u := args[0]

	keys, _ := ioutil.ReadDir("/git/keys/" + u)

	for _, k := range keys {
		if k.IsDir() {
			continue
		}

		key, _ := ioutil.ReadFile("/git/keys/" + u + "/" + k.Name())

		fmt.Println(string(key))
	}
}
