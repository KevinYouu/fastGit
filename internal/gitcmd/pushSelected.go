package gitcmd

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/logs"
	"github.com/KevinYouu/fastGit/internal/theme"
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

	// 显示开始信息
	fmt.Printf("%s",
		theme.TitleStyle.Render("Starting Git Push Selected Process"))

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

	// 使用新的命令执行器执行Git操作
	commands := []command.CommandInfo{
		{
			Command:     "git",
			Args:        append([]string{"add"}, data...),
			Description: "Adding selected files to staging area",
			LoadingMsg:  "Adding selected files...",
			SuccessMsg:  "Selected files added successfully",
		},
		{
			Command:     "git",
			Args:        []string{"commit", "-m", commitMessage},
			Description: "Creating commit with message",
			LoadingMsg:  "Creating commit...",
			SuccessMsg:  "Commit created successfully",
		},
		{
			Command:     "git",
			Args:        []string{"pull"},
			Description: "Pulling latest changes from remote",
			LoadingMsg:  "Pulling changes...",
			SuccessMsg:  "Pull completed successfully",
		},
		{
			Command:     "git",
			Args:        []string{"push"},
			Description: "Pushing changes to remote repository",
			LoadingMsg:  "Pushing to remote...",
			SuccessMsg:  "Push completed successfully",
		},
	}

	return command.RunMultipleCommands(commands)
}
