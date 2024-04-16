package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

func PushSelected() {
	fileStatuss, err := status.GetFileStatuses()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(1)
	}
	if len(fileStatuss) == 0 {
		fmt.Println(colors.RenderColor("blue", "No files to push."))
		os.Exit(0)
	}

	var selectedFiles []string
	for _, fileStatus := range fileStatuss {
		if fileStatus.Status != "" {
			selectedFiles = append(selectedFiles, fileStatus.Path)
		}
	}

	// data := multipleChoice.MultipleChoice(selectedFiles)
	data, err := form.MultiSelectForm("Select files to push", selectedFiles)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(0)
		return
	}

	if len(data) == 0 {
		fmt.Println(colors.RenderColor("red", "No files selected."))
		os.Exit(0)
	}
	// suffix := choose.Choose([]string{"fix", "feat", "refactor", "style", "chore", "docs", "test", "revert"})
	// commitMessage := input.Input("Enter your commit message: ", "commit message", "(esc to quit)", suffix+": ")
	options := []form.Option{
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

	addLog, err := command.RunCommand("git", append([]string{"add"}, data...)...)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to add files: "+err.Error()))
		return
	}
	fmt.Println(addLog, colors.RenderColor("green", "Files added successfully.\n"))

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
}
