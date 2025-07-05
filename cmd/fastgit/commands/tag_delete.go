package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func TagDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "td",
		Aliases: []string{"tag-delete", "deltag"},
		Short:   "Delete a tag from local and remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.DeleteAndPushTag(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
