package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/spf13/cobra"
)

func PushSelectedCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "push-selected",
		Aliases: []string{"ps"},
		Short:   "Push selected changes",
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.PushSelected(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
