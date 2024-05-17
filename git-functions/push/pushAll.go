package push

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/config"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/logs"
	"github.com/KevinYouu/fastGit/functions/spinner"
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

	spinner.Spinner("Pushing...", "done", func() {
		addlog, err := command.RunCommand("git", "add", "-A")
		if err != nil {
			 logs.Error("Failed to add files: ")
			 fmt.Println(addlog)
			return
		}
		logs.Success("Files added successfully.\n")

		commLog, err := command.RunCommand("git", "commit", "-m", commitMessage)
		if err != nil {
			logs.Error("Failed to commit: ")
			fmt.Println(err.Error())
			return
		}
		logs.Success("Commit successful.\n")
		fmt.Println(commLog)

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
		 logs.Success("Push successful.\n")
	})
}
