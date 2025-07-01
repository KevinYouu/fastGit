package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func MergeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "merge",
		Aliases: []string{"m"},
		Short:   "Merge into current branch",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.MergeIntoCurrent(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
