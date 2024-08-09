package main

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/pkg/components/version"
	"github.com/KevinYouu/fastGit/pkg/gitcmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fastGit",
	Short: "fastGit is a tool that helps you quickly submit code with a command line interface.",
	Run: func(cmd *cobra.Command, args []string) {
		version.GetVersion()
	},
}

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "c",
			Short: "Clone a repository",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.Clone()
			},
		},
		&cobra.Command{
			Use:   "pa",
			Short: "Push all changes",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.PushAll()
			},
		},
		&cobra.Command{
			Use:   "ps",
			Short: "Push selected changes",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.PushSelected()
			},
		},
		&cobra.Command{
			Use:   "ra",
			Short: "Add a remote",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.AddRemote()
			},
		},
		&cobra.Command{
			Use:   "rv",
			Short: "View remotes",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.GetRemotes()
			},
		},
		&cobra.Command{
			Use:   "rs",
			Short: "Reset changes",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.Reset()
			},
		},
		&cobra.Command{
			Use:   "t",
			Short: "Create and push a tag",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.CreateAndPushTag()
			},
		},
		&cobra.Command{
			Use:   "s",
			Short: "Show status",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.Status()
			},
		},
		&cobra.Command{
			Use:   "m",
			Short: "Merge into current branch",
			Run: func(cmd *cobra.Command, args []string) {
				gitcmd.MergeIntoCurrent()
			},
		},
		&cobra.Command{
			Use:   "v",
			Short: "Show version",
			Run: func(cmd *cobra.Command, args []string) {
				version.GetVersion()
			},
		},
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
