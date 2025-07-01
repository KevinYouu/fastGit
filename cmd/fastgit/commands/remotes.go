package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func RemotesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "remotes",
		Aliases: []string{"rv"},
		Short:   "View git remotes",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.GetRemotes(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
