package gitcmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/KevinYouu/fastGit/pkg/components/colors"
	"github.com/KevinYouu/fastGit/pkg/components/form"
)

// create a new tag and push it to the remote repository.
func CreateAndPushTag() {
	latestVersion, err := GetLatestTag()
	if err != nil {
		log.Printf("get latest tag error: %s", err)
		return
	}
	newVersion := incrementVersion(latestVersion)

	version, err := form.Input("Enter your version: ", newVersion)
	if err != nil {
		log.Printf("get version error: %s", err)
		return
	}

	commitMessage, err := form.Input("Enter your commit message: ", "")
	if err != nil {
		log.Printf("get commit message error: %s", err)
		return
	}

	cmd := exec.Command("git", "tag", "-a", version, "-m", commitMessage)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), colors.RenderColor("red", "Failed to create tag: "+string(output)))
		return
	}
	fmt.Println(string(output), colors.RenderColor("green", "Tag created successfully."))

	cmd = exec.Command("git", "push", "origin", newVersion)
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output), colors.RenderColor("red", "Failed to push tag: "+string(output)))
		return
	}
	fmt.Println(string(output), colors.RenderColor("green", "Tag pushed successfully."))
}

// get the latest tag from the repository.
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

// increment the version number.
func incrementVersion(currentVersion string) string {
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(currentVersion)
	if len(matches) != 4 {
		fmt.Println("Invalid version format:", currentVersion)
		os.Exit(1)
	}

	// parse the version number
	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	// increment the patch number
	patch++
	if patch > 9 {
		patch = 0
		minor++
		if minor > 9 {
			minor = 0
			major++
			// if major > 99 {
			// 	fmt.Println("Version number out of range")
			// 	os.Exit(1)
			// }
		}
	}

	//
	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	return newVersion
}
