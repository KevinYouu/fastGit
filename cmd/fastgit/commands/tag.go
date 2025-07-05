package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/spf13/cobra"
)

func TagCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tag",
		Aliases: []string{"t"},
		Short:   i18n.T("tag.short"),
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.CreateAndPushTag(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
