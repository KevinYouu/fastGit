package tag

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/input"
)

// CreateAndPushTag 创建标签并推送到远程仓库
func CreateAndPushTag() {
	commitMessage := input.Input("Enter your commit message: \n", "tag commit message", "\n(esc to quit)")

	latestVersion, err := GetLatestTag()
	if err != nil {
		log.Printf("get latest tag error: %s", err)
		return
	}
	newVersion := incrementVersion(latestVersion)

	cmd := exec.Command("git", "tag", "-a", newVersion, "-m", commitMessage)
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

// Basic example of how to list tags.
// GetLatestTag 获取最新的标签版本号，如果没有标签则返回 "0.0.0"
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
	// 解析当前版本号
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	matches := re.FindStringSubmatch(currentVersion)
	if len(matches) != 4 {
		fmt.Println("Invalid version format:", currentVersion)
		os.Exit(1)
	}

	// 将版本号的每个部分转换为整数
	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	// 版本号自增逻辑
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

	// 重新构建版本号字符串
	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	return newVersion
}
