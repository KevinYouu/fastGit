package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func StatusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "status",
		Aliases: []string{"s"},
		Short:   "Show git status",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.Status(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
