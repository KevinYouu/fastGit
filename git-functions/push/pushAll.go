package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/input"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

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
	suffix := choose.Choose([]string{"fix", "feat", "refactor", "style", "chore", "docs", "test", "revert"})
	commitMessage := input.Input("Enter your commit message: ", "commit message", "(esc to quit)", suffix+": ")

	log, err := command.RunCommand("git", "pull")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to pull: "+err.Error()))
		return
	} else {
		fmt.Println(log, colors.RenderColor("green", "Pulled successfully.\n"))
	}

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

	pushLog, err := command.RunCommand("git", "push")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to push: "+err.Error()))
		return
	}
	fmt.Println(pushLog, colors.RenderColor("green", "Push successful."))
}
