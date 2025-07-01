package gitcmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/KevinYouu/fastGit/internal/colors"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
)

func CreateAndPushTag() error {
	latestVersion, err := GetLatestTag()
	if err != nil {
		return fmt.Errorf("get latest tag error: %w", err)
	}
	newVersion := incrementVersion(latestVersion)

	version, err := form.Input("Enter your version: ", newVersion)
	if err != nil {
		return fmt.Errorf("get version error: %w", err)
	}

	commitMessage, err := form.Input("Enter your commit message: ", "")
	if err != nil {
		return fmt.Errorf("get commit message error: %w", err)
	}

	cmd := exec.Command("git", "tag", "-a", version, "-m", commitMessage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to create tag: %s", string(output))
	}
	fmt.Println(string(output), colors.RenderColor("green", "Tag created successfully."))

	cmd = exec.Command("git", "push", "origin", version)
	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to push tag: %s", string(output))
	}
	fmt.Println(string(output), colors.RenderColor("green", "Tag pushed successfully."))
	return nil
}

func GetLatestTag() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	output, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok && exitError.ExitCode() == 128 {
			return "0.0.0", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func incrementVersion(currentVersion string) string {
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(currentVersion)
	if len(matches) != 4 {
		return "0.0.0"
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	maxPatch, err := config.GetTagPatch()
	if err != nil {
		return "0.0.0"
	}

	patch++
	if patch > maxPatch.Patch {
		patch = 0
		minor++
		if minor > maxPatch.Minor {
			minor = 0
			major++
		}
	}

	newVersion := fmt.Sprintf("%s%d.%d.%d%s", maxPatch.Prefix, major, minor, patch, maxPatch.Suffix)
	return newVersion
}
