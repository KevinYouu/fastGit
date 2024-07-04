package merge

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/form"
	"github.com/KevinYouu/fastGit/functions/logs"
)

func MergeIntoCurrent() {
	branches, err := getCurrentBranches()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	selectedBranch, err := selectBranchToMerge(branches)
	if err != nil {
		fmt.Println(err)
		return
	}

	output, err := command.RunCmd("git", []string{"merge", selectedBranch}, "Merge branch successfully.")
	if err != nil {
		logs.Error("Failed to merge: " + output)
		return
	}
}

func getCurrentBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running git branch: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var branches []string

	for _, line := range lines {
		branch := strings.TrimSpace(strings.TrimPrefix(line, "* "))
		if branch != "" && branch != "(no branch)" {
			branches = append(branches, branch)
		}
	}

	return branches, nil
}

func selectBranchToMerge(branches []string) (string, error) {
	_, selectedBranch, err := form.SelectFormWithStringSlice("Branch name to merge into the current branch", branches)
	if err != nil {
		return "", fmt.Errorf("error selecting branch: %v", err)
	}
	return selectedBranch, nil
}
