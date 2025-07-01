package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func PushAllCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "push-all",
		Aliases: []string{"pa"},
		Short:   "Push all changes",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.PushAll(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
