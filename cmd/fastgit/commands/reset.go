package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func ResetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "reset",
		Aliases: []string{"rs"},
		Short:   "Reset to a specific commit",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.Reset(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
