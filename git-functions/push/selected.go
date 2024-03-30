package push

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/KevinYouu/fastGit/functions/choose"
	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"
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

	cmd := exec.Command("git", "pull")
	err = cmd.Run()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to pull: "+err.Error()))
		os.Exit(1)
	} else {
		fmt.Println(colors.RenderColor("green", "Pulled successfully."))
	}

	suffix := choose.Choose([]string{"fix", "feat", "docs", "style", "refactor", "test", "chore", "revert"})
	commitMessage := input.Input("Enter your commit message: \n", "commit message", "\n(esc to quit)")

	cmd = exec.Command("git", append([]string{"add"}, data...)...)
	err = cmd.Run()
	if err != nil {
		// fmt.Println("Error executing git add command:", err)
		fmt.Println(colors.RenderColor("red", "Failed to add files: "+err.Error()))
		return
	}
	fmt.Println(colors.RenderColor("green", "Files added successfully."))

	cmd = exec.Command("git", "commit", "-m", suffix+" "+commitMessage)
	err = cmd.Run()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to commit: "+err.Error()))
		return
	}

	fmt.Println(colors.RenderColor("green", "Commit successful."))

	// 执行 git push 命令
	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to push: "+err.Error()))
		return
	}

	fmt.Println(colors.RenderColor("green", "Push successful."))
}
