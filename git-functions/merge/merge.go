package merge

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/form"
)

func MergeIntoCurrent() {
	cmd := exec.Command("git", "branch")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	lines := strings.Split(string(output), "\n")

	branches := make([]string, 0)

	for _, line := range lines {
		branch := strings.TrimSpace(strings.TrimPrefix(line, "* "))
		if branch != "" && branch != "(no branch)" {
			branches = append(branches, branch)
		}
	}

	_, value, err := form.SelectFormWithStringSlice("Branch name to merge into the current branch", branches)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mergeLog, err := command.RunCommand("git", "merge", value)
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to commit: "+err.Error()))
		return
	}

	fmt.Println(mergeLog)
	fmt.Println(colors.RenderColor("green", "Merge branch successfully"))
}
