package push

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/multipleChoice"
	"github.com/KevinYouu/fastGit/git-functions/status"
)

func PushSelected() {
	fileStatuss, err := status.GetFileStatuses()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(1)
	}

	var selectedFiles []string
	for _, fileStatus := range fileStatuss {
		if fileStatus.Status != "" {
			selectedFiles = append(selectedFiles, fileStatus.Path)
		}
	}

	data := multipleChoice.MultipleChoice(selectedFiles)

	cmd := exec.Command("git", append([]string{"add"}, data...)...)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing git add command:", err)
		return
	}

	fmt.Println("Files added successfully.")
}
