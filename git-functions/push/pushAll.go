package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/spinner"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

var options = []form.Option{
	{Label: "fix", Value: "fix"},
	{Label: "feat", Value: "feat"},
	{Label: "refactor", Value: "refactor"},
	{Label: "chore", Value: "chore"},
	{Label: "build", Value: "build"},
	{Label: "revert", Value: "revert"},
	{Label: "style", Value: "style"},
	{Label: "docs", Value: "docs"},
	{Label: "test", Value: "test"},
}

func PushAll() {
	fileStatuss, err := status.GetFileStatuses()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(1)
	}
	if len(fileStatuss) == 0 {
		fmt.Println(colors.RenderColor("blue", "No files to push."))
		os.Exit(0)
	}
	// suffix := choose.Choose([]string{"fix", "feat", "refactor", "style", "chore", "docs", "test", "revert"})
	_, suffix, err := form.SelectForm(options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// commitMessage := input.Input("Enter your commit message: ", "commit message", "(esc to quit)", suffix+": ")
	commitMessage, err := form.Input("Enter your commit message: ", suffix+": ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	spinner.Spinner("Pushing...", "done", func() {
		addlog, err := command.RunCommand("git", "add", "-A")
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to add files: "+err.Error()))
			return
		}
		fmt.Println(addlog, colors.RenderColor("green", "Files added successfully.\n"))

		commLog, err := command.RunCommand("git", "commit", "-m", commitMessage)
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to commit: "+err.Error()))
			return
		}
		fmt.Println(commLog, colors.RenderColor("green", "Commit successful.\n"))

		pullLog, err := command.RunCommand("git", "pull")
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to pull: "+err.Error()))
			return
		} else {
			fmt.Println(pullLog, colors.RenderColor("green", "Pulled successfully.\n"))
		}

		pushLog, err := command.RunCommand("git", "push")
		if err != nil {
			fmt.Println(colors.RenderColor("red", "Failed to push: "+err.Error()))
			return
		}
		fmt.Println(pushLog, colors.RenderColor("green", "Push successful."))
	})
}
