package commands

import (
	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/KevinYouu/fastGit/internal/version"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   i18n.T("version.short"),
		Run: func(cmd *cobra.Command, args []string) {
			version.GetVersion()
		},
	}
	return cmd
}
