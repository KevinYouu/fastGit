package commands

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/gitcmd"
	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/spf13/cobra"
)

func CherryPickCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "cherry-pick",
		Aliases: []string{"cp"},
		Short:   i18n.T("cherry.pick.short"),
		Run: func(cmd *cobra.Command, args []string) {
			if err := gitcmd.CherryPick(); err != nil {
				fmt.Println(err)
			}
		},
	}
	return cmd
}
