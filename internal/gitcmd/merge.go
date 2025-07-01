package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/logs"
)

func MergeIntoCurrent() error {
	branches, err := getCurrentBranches()
	if err != nil {
		return fmt.Errorf("getCurrentBranches: %w", err)
	}

	if len(branches) == 0 {
		logs.Info("No branches to merge.")
		return nil
	}

	selectedBranch, err := selectBranchToMerge(branches)
	if err != nil {
		return fmt.Errorf("selectBranchToMerge: %w", err)
	}

	output, err := command.RunCmd("git", []string{"merge", selectedBranch}, "Merge branch successfully.")
	if err != nil {
		logs.Error("Failed to merge: " + output)
		return fmt.Errorf("git merge: %s", output)
	}
	return nil
}

func getCurrentBranches() ([]string, error) {
	cmd := exec.Command("git", "branch")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("error running git branch: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var branches []string

	currentBranch, err := exec.Command("git", "branch", "--show-current").CombinedOutput()
	if err != nil {
		logs.Error("Failed to get current branch: " + string(currentBranch))
		return nil, err
	}

	current := strings.TrimSpace(string(currentBranch))

	for _, line := range lines {
		branch := strings.TrimSpace(strings.TrimPrefix(line, "* "))
		if branch != "" && branch != "(no branch)" && branch != current {
			branches = append(branches, branch)
		}
	}

	return branches, nil
}

func selectBranchToMerge(branches []string) (string, error) {
	_, selectedBranch, err := form.SelectFormWithStringSlice("Branch name to merge into the current branch", branches)
	if err != nil {
		return "", fmt.Errorf("error selecting branch: %w", err)
	}
	return selectedBranch, nil
}
