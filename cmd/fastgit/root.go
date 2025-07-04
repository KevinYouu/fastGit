package main

import (
	"github.com/KevinYouu/fastGit/cmd/fastgit/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fastGit",
	Short: "fastGit is a tool that helps you quickly submit code with a command line interface.",
	Run: func(cmd *cobra.Command, args []string) {
		commands.VersionCommand().Run(cmd, args)
	},
}

func init() {
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
