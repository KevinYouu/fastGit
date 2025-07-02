package gitcmd

import (
	"fmt"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/logs"
)

func PushAll() error {
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
			Args:        []string{"add", "-A"},
			Description: "Adding all files to staging area",
			LoadingMsg:  "Adding files...",
			SuccessMsg:  "Files added successfully",
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
