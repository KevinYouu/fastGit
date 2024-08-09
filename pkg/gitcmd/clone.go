package gitcmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/KevinYouu/fastGit/pkg/components/form"
	"github.com/KevinYouu/fastGit/pkg/components/logs"
)

func Clone() {
	cloneURL, err := form.Input("Enter the URL of the repository you want to clone: ", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd := exec.Command("git", "clone", cloneURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logs.Error("Failed to clone: \n" + string(output))
		return
	}
	fmt.Println(string(output))
	logs.Success("Cloned successfully.")
	folderName := filepath.Base(strings.TrimSuffix(cloneURL, ".git"))

	confirm := form.Confirm("Open repository in vscode?")
	if confirm {
		cmd = exec.Command("code", folderName)
		cmd.Run()
	}
}
