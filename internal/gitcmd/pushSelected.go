package gitcmd

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/logs"
)

func PushSelected() error {
	fileStatus, err := getFileStatuses()
	if err != nil {
		logs.Error("Failed to get file statuses")
		return fmt.Errorf("getFileStatuses: %w", err)
	}
	if len(fileStatus) == 0 {
		logs.Info("No files to push.")
		return nil
	}

	var selectedFiles []string
	for _, fileStatus := range fileStatus {
		if fileStatus.Status != "" {
			selectedFiles = append(selectedFiles, fileStatus.Path)
		}
	}

	data, err := form.MultiSelectForm("Select files to push", selectedFiles)
	if err != nil {
		logs.Error("Failed to get file statuses:")
		return fmt.Errorf("MultiSelectForm: %w", err)
	}

	if len(data) == 0 {
		logs.Error("No files selected.")
		return nil
	}
	options, err := config.GetOptions()
	if err != nil {
		logs.Error("Failed to get options:")
		return fmt.Errorf("GetOptions: %w", err)
	}

	_, suffix, err := form.SelectForm("Choose a commit type", options)
	if err != nil {
		return fmt.Errorf("SelectForm: %w", err)
	}
	config.IncrementUsage(suffix)

	commitMessage, err := form.Input("Enter your commit message: ", suffix+": ")
	if err != nil {
		return fmt.Errorf("Input: %w", err)
	}

	output, err := command.RunCmd("git", append([]string{"add"}, data...), "Added files successfully.")
	if err != nil {
		logs.Error("Failed to add files: " + output)
		return fmt.Errorf("git add: %s", output)
	}

	output, err = command.RunCmd("git", []string{"commit", "-m", commitMessage}, "Commit successfully.")
	if err != nil {
		logs.Error("Failed to commit: " + output)
		return fmt.Errorf("git commit: %s", output)
	}

	output, err = command.RunCmd("git", []string{"pull"}, "Pulled successfully.")
	if err != nil {
		logs.Error("Failed to pull: " + output)
		return fmt.Errorf("git pull: %s", output)
	}

	output, err = command.RunCmd("git", []string{"push"}, "Pushed successfully.")
	if err != nil {
		logs.Error("Failed to push: " + output)
		return fmt.Errorf("git push: %s", output)
	}
	return nil
}
