package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/KevinYouu/fastGit/internal/logs"
)

// CherryPickOption represents different cherry-pick options
type CherryPickOption struct {
	Name           string
	Args           []string
	NameKey        string // i18n key for option name
	DescriptionKey string // i18n key for option description
}

var cherryPickOptions = []CherryPickOption{
	{
		Name:           "default",
		Args:           []string{},
		NameKey:        "cherry.pick.option.default.name",
		DescriptionKey: "cherry.pick.option.default.description",
	},
	{
		Name:           "no-commit",
		Args:           []string{"--no-commit"},
		NameKey:        "cherry.pick.option.no.commit.name",
		DescriptionKey: "cherry.pick.option.no.commit.description",
	},
	{
		Name:           "edit",
		Args:           []string{"--edit"},
		NameKey:        "cherry.pick.option.edit.name",
		DescriptionKey: "cherry.pick.option.edit.description",
	},
	{
		Name:           "signoff",
		Args:           []string{"--signoff"},
		NameKey:        "cherry.pick.option.signoff.name",
		DescriptionKey: "cherry.pick.option.signoff.description",
	},
}

func CherryPick() error {
	// First check if we're in a git repository
	if !isGitRepository() {
		return fmt.Errorf("%s", i18n.T("error.not.git.repo"))
	}

	// Get all commits from all branches
	commits, err := getAllCommitsForCherryPick()
	if err != nil {
		return fmt.Errorf(i18n.T("cherry.pick.error.get.commits")+": %v", err)
	}

	if len(commits) == 0 {
		return fmt.Errorf("%s", i18n.T("cherry.pick.error.no.commits"))
	}

	// Let user select commits to cherry-pick
	selectedCommits, err := selectCommitsForCherryPick(commits)
	if err != nil {
		return err
	}

	if len(selectedCommits) == 0 {
		logs.Info(i18n.T("cherry.pick.no.commits.selected"))
		return nil
	}

	// Let user select cherry-pick options
	option, err := selectCherryPickOption()
	if err != nil {
		return err
	}

	// Execute cherry-pick for each selected commit
	for _, commit := range selectedCommits {
		if err := executeCherryPick(commit, option); err != nil {
			return fmt.Errorf(i18n.T("cherry.pick.error.execute")+": %v", err)
		}
		logs.Success(i18n.T("cherry.pick.success.commit") + ": " + commit.Hash[:8])
	}

	logs.Success(i18n.T("cherry.pick.success.all"))
	return nil
}

func getAllCommitsForCherryPick() ([]Commit, error) {
	// Get all branches
	cmd := exec.Command("git", "branch", "-a", "--format=%(refname:short)")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	branches := strings.Fields(string(output))
	var allCommits []Commit

	// Get commits from each branch
	for _, branch := range branches {
		// Skip remote tracking branches that are duplicates
		if strings.HasPrefix(branch, "origin/") {
			continue
		}

		cmd := exec.Command("git", "log", branch, "--pretty=format:%H|%s|%an|%ae|%ad", "--date=short", "-20")
		output, err := cmd.Output()
		if err != nil {
			continue // Skip if branch doesn't exist or has no commits
		}

		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		for _, line := range lines {
			if line == "" {
				continue
			}
			parts := strings.Split(line, "|")
			if len(parts) >= 5 {
				commit := Commit{
					Hash:    parts[0],
					Message: parts[1],
					Author:  parts[2],
					Email:   parts[3],
					Date:    parts[4],
					IsHead:  false,
				}
				allCommits = append(allCommits, commit)
			}
		}
	}

	return allCommits, nil
}

func selectCommitsForCherryPick(commits []Commit) ([]Commit, error) {
	// Create options for the multi-select form
	var options []string
	for _, commit := range commits {
		displayText := fmt.Sprintf("[%s] %s (%s) - %s",
			commit.Hash[:8],
			commit.Date,
			commit.Author,
			commit.Message)

		options = append(options, displayText)
	}

	// Show multi-select form
	selectedDisplayTexts, err := form.MultiSelectForm(
		i18n.T("cherry.pick.select.commits"),
		options,
	)
	if err != nil {
		return nil, err
	}

	// Convert selected display texts back to commits
	var selectedCommits []Commit
	for _, selectedText := range selectedDisplayTexts {
		for i, option := range options {
			if option == selectedText && i < len(commits) {
				selectedCommits = append(selectedCommits, commits[i])
				break
			}
		}
	}

	return selectedCommits, nil
}

func selectCherryPickOption() (CherryPickOption, error) {
	var options []config.Option
	for _, option := range cherryPickOptions {
		options = append(options, config.Option{
			Label: i18n.T(option.NameKey),
			Value: option.Name,
		})
	}

	_, selectedName, err := form.SelectForm(
		i18n.T("cherry.pick.select.option"),
		options,
	)
	if err != nil {
		return CherryPickOption{}, err
	}

	for _, option := range cherryPickOptions {
		if option.Name == selectedName {
			return option, nil
		}
	}

	return cherryPickOptions[0], nil // fallback to default
}

func executeCherryPick(commit Commit, option CherryPickOption) error {
	// Build the git cherry-pick command
	args := []string{"cherry-pick"}
	args = append(args, option.Args...)
	args = append(args, commit.Hash)

	// Show what we're about to do
	cmdStr := "git " + strings.Join(args, " ")
	logs.Info(i18n.T("cherry.pick.executing") + ": " + cmdStr)

	// Execute the command and get detailed error information
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		// Check for specific cherry-pick errors
		outputStr := string(output)

		if strings.Contains(outputStr, "CONFLICT") {
			logs.Error(i18n.T("cherry.pick.conflict.detected"))
			logs.Info(i18n.T("cherry.pick.conflict.instructions"))
			logs.Info(i18n.T("cherry.pick.conflict.output") + ":\n" + outputStr)
			return fmt.Errorf("%s", i18n.T("cherry.pick.conflict.resolution.needed"))
		}

		if strings.Contains(outputStr, "empty commit") {
			logs.Waring(i18n.T("cherry.pick.empty.commit") + ": " + commit.Hash[:8])
			// For empty commits, we might want to continue with --allow-empty or skip
			return fmt.Errorf("%s", i18n.T("cherry.pick.empty.commit.error"))
		}

		if strings.Contains(outputStr, "already exists") || strings.Contains(outputStr, "nothing to commit") {
			logs.Waring(i18n.T("cherry.pick.already.applied") + ": " + commit.Hash[:8])
			return nil // Skip this commit as it's already applied
		}

		// Generic error with output
		logs.Error(i18n.T("cherry.pick.failed.output") + ":\n" + outputStr)
		return fmt.Errorf(i18n.T("cherry.pick.failed.generic")+": %v", err)
	}

	// Success case
	if strings.TrimSpace(string(output)) != "" {
		logs.Info(string(output))
	}

	return nil
}

func isGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run() == nil
}
