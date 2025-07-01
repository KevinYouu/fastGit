package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func TagCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tag",
		Aliases: []string{"t"},
		Short:   "Create and push a tag",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.CreateAndPushTag(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
