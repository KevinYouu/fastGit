package clone

import (
	"fmt"
	"os/exec"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"
)

func Clone() {
	cloneURL := input.Input("Enter the URL of the repository you want to clone: ", "clone url", "(esc to quit)", "")

	cmd := exec.Command("git", "clone", cloneURL)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), colors.RenderColor("red", "Failed to clone: "+string(output)))
		return
	}
	fmt.Println(string(output), colors.RenderColor("green", "Files added successfully."))
}
