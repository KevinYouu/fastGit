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

func PushAll() {
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

	output, err := command.RunCmd("git", []string{"add", "-A"}, "Files added successfully.")
	if err != nil {
		logs.Error("Failed to add: " + output)
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
