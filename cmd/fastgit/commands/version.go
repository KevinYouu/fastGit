package commands

import (
	"github.com/KevinYouu/fastGit/internal/version"
	"github.com/spf13/cobra"
)

func VersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Show fastGit version",
		Run: func(cmd *cobra.Command, args []string) {
			version.GetVersion()
		},
	}
	return cmd
}
