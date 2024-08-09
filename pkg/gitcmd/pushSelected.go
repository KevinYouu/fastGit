package gitcmd

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/pkg/components/command"
	"github.com/KevinYouu/fastGit/pkg/components/config"
	"github.com/KevinYouu/fastGit/pkg/components/form"
	"github.com/KevinYouu/fastGit/pkg/components/logs"
)

func PushSelected() {
	fileStatus, err := getFileStatuses()
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

	output, err := command.RunCmd("git", append([]string{"add"}, data...), "Added files successfully.")
	if err != nil {
		logs.Error("Failed to add files: " + output)
		return
	}

	output, err = command.RunCmd("git", []string{"commit", "-m", commitMessage}, "Commit successfully.")
	if err != nil {
		logs.Error("Failed to commit: " + output)
		return
	}

	output, err = command.RunCmd("git", []string{"pull"}, "Pulled successfully.")
	if err != nil {
		logs.Error("Failed to pull: " + output)
		return
	}

	output, err = command.RunCmd("git", []string{"push"}, "Pushed successfully.")
	if err != nil {
		logs.Error("Failed to push: " + output)
		return
	}
}
