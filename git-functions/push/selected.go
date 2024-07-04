package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/config"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/logs"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

func PushSelected() {
	fileStatus, err := status.GetFileStatuses()
	if err != nil {
		fmt.Println(err)
		logs.Error("Failed to get file statuses")
		os.Exit(1)
	}
	if len(fileStatus) == 0 {
		logs.Info("No files to push.")
		os.Exit(0)
	}

	var selectedFiles []string
	for _, fileStatus := range fileStatus {
		if fileStatus.Status != "" {
			selectedFiles = append(selectedFiles, fileStatus.Path)
		}
	}

	// data := multipleChoice.MultipleChoice(selectedFiles)
	data, err := form.MultiSelectForm("Select files to push", selectedFiles)
	if err != nil {
		logs.Error("Failed to get file statuses:")
		fmt.Println(err)
		os.Exit(0)
		return
	}

	if len(data) == 0 {
		logs.Error("No files selected.")
		os.Exit(0)
	}
	options, err := config.GetOptions()
	if err != nil {
		logs.Error("Failed to get options:")
		fmt.Println(err)
		os.Exit(1)
	}

	_, suffix, err := form.SelectForm("Choose a commit type", options)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	commitMessage, err := form.Input("Enter your commit message: ", suffix+": ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addLog, err := command.RunCommand("git", append([]string{"add"}, data...)...)
	if err != nil {
		logs.Error("Failed to add files: ")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(addLog)
	logs.Success("Files added successfully.\n")

	commLog, err := command.RunCommand("git", "commit", "-m", commitMessage)
	if err != nil {
		logs.Error("Failed to commit: ")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(commLog)
	logs.Success("Commit successful.\n")

	pullLog, err := command.RunCommand("git", "pull")
	if err != nil {
		logs.Error("Failed to pull: ")
		fmt.Println(err.Error())
		return
	} else {
		fmt.Println(pullLog)
		logs.Success("Pulled successfully.\n")
	}

	pushLog, err := command.RunCommand("git", "push")
	if err != nil {
		logs.Error("Failed to push: ")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(pushLog)
	logs.Success("Pushed successfully.")
}
