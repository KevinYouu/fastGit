package main

import (
	"github.com/KevinYouu/fastGit/cmd/fastgit/commands"
	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fastGit",
	Short: "", // Will be set dynamically
	Run: func(cmd *cobra.Command, args []string) {
		commands.VersionCommand().Run(cmd, args)
	},
}

// updateRootCommandDescriptions updates all command descriptions based on current language
func updateRootCommandDescriptions() {
	rootCmd.Short = i18n.T("root.short")

	// Update all subcommands
	for _, cmd := range rootCmd.Commands() {
		switch cmd.Use {
		case "version":
			cmd.Short = i18n.T("version.short")
		case "status":
			cmd.Short = i18n.T("status.short")
		case "push-all":
			cmd.Short = i18n.T("push.all.short")
		case "push-selected":
			cmd.Short = i18n.T("push.selected.short")
		case "remotes":
			cmd.Short = i18n.T("remotes.short")
		case "reset":
			cmd.Short = i18n.T("reset.short")
		case "tag":
			cmd.Short = i18n.T("tag.short")
		case "td":
			cmd.Short = i18n.T("tag.delete.short")
		case "merge":
			cmd.Short = i18n.T("merge.short")
		case "update":
			cmd.Short = i18n.T("update.short")
		case "init":
			cmd.Short = i18n.T("init.short")
		}
	}
}

func init() {
	// Add language flag support
	i18n.AddLanguageFlag(rootCmd)

	rootCmd.AddCommand(
		commands.PushAllCommand(),
		commands.PushSelectedCommand(),
		commands.RemotesCommand(),
		commands.ResetCommand(),
		commands.TagCommand(),
		commands.TagDeleteCommand(),
		commands.StatusCommand(),
		commands.MergeCommand(),
		commands.VersionCommand(),
		commands.UpdateCommand(),
		commands.InitCommand(),
	)
}
