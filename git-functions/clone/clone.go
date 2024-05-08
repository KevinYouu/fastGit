package clone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/spinner"
)

func Clone() {
	cloneURL, err := form.Input("Enter the URL of the repository you want to clone: ", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cmd := exec.Command("git", "clone", cloneURL)
	spinner.Spinner("Cloning repository...", "done", func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(string(output), colors.RenderColor("red", "Failed to clone: "+string(output)))
			return
		}
		fmt.Println(string(output), colors.RenderColor("green", "Cloned successfully."))
		// Extract folder name from repo URL
	})
	folderName := filepath.Base(strings.TrimSuffix(cloneURL, ".git"))

	confirm := form.Confirm("Open repository in vscode?")
	if confirm {
		cmd = exec.Command("code", folderName)
		cmd.Run()
	}
}
