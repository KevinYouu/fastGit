package gitcmd

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/i18n"
)

func CreateAndPushTag() error {
	latestVersion, err := GetLatestTag()
	if err != nil {
		return fmt.Errorf("get latest tag error: %w", err)
	}
	newVersion := incrementVersion(latestVersion)

	version, err := form.Input(i18n.T("tag.input.version"), newVersion)
	if err != nil {
		return fmt.Errorf("get version error: %w", err)
	}

	commitMessage, err := form.Input(i18n.T("tag.input.commit.message"), "")
	if err != nil {
		return fmt.Errorf("get commit message error: %w", err)
	}

	// 使用新的命令执行器
	commands := []command.CommandInfo{
		{
			Command:     "git",
			Args:        []string{"tag", "-a", version, "-m", commitMessage},
			Description: i18n.T("tag.create.description"),
			LoadingMsg:  i18n.T("tag.create.loading"),
			SuccessMsg:  fmt.Sprintf(i18n.T("tag.create.success"), version),
		},
		{
			Command:     "git",
			Args:        []string{"push", "origin", version},
			Description: i18n.T("tag.push.description"),
			LoadingMsg:  i18n.T("tag.push.loading"),
			SuccessMsg:  fmt.Sprintf(i18n.T("tag.push.success"), version),
		},
	}

	return command.RunMultipleCommands(commands)
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

// GetAllTags 获取所有标签列表，按创建时间排序（最新的在前）
func GetAllTags() ([]string, error) {
	// 使用 git tag --sort=-creatordate 来按创建时间倒序排列
	cmd := exec.Command("git", "tag", "--sort=-creatordate")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get tags: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	var tags []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			tags = append(tags, line)
		}
	}

	if len(tags) == 0 {
		return nil, fmt.Errorf("%s", i18n.T("tag.no.tags"))
	}

	return tags, nil
}

// GetTagWithCreationDate 获取标签的创建时间
func GetTagWithCreationDate(tag string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--pretty=format:%ad", "--date=format:%Y-%m-%d %H:%M", tag)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// DeleteAndPushTag 删除标签并推送到远程仓库
func DeleteAndPushTag() error {
	// 获取所有标签
	tags, err := GetAllTags()
	if err != nil {
		return fmt.Errorf("get tags error: %w", err)
	}

	// 将标签转换为选项格式，包含创建时间
	var options []config.Option
	for _, tag := range tags {
		creationDate, err := GetTagWithCreationDate(tag)
		if err != nil {
			// 如果获取时间失败，只显示标签名
			options = append(options, config.Option{
				Label: fmt.Sprintf("🏷️  %s", tag),
				Value: tag,
			})
		} else {
			// 显示标签名和创建时间
			options = append(options, config.Option{
				Label: fmt.Sprintf("🏷️  %s  📅 %s", tag, creationDate),
				Value: tag,
			})
		}
	}

	// 让用户选择要删除的标签
	_, selectedTag, err := form.SelectForm(i18n.T("tag.delete.select"), options)
	if err != nil {
		return fmt.Errorf("select tag error: %w", err)
	}

	// 确认删除操作
	confirmMessage := fmt.Sprintf(i18n.T("tag.delete.confirm"), selectedTag)
	if !form.Confirm(confirmMessage) {
		fmt.Println(i18n.T("tag.delete.cancelled"))
		return nil
	}

	// 执行删除操作
	commands := []command.CommandInfo{
		{
			Command:     "git",
			Args:        []string{"tag", "-d", selectedTag},
			Description: i18n.T("tag.delete.local"),
			LoadingMsg:  fmt.Sprintf(i18n.T("tag.delete.local.loading"), selectedTag),
			SuccessMsg:  fmt.Sprintf(i18n.T("tag.delete.local.success"), selectedTag),
		},
		{
			Command:     "git",
			Args:        []string{"push", "origin", ":refs/tags/" + selectedTag},
			Description: i18n.T("tag.delete.remote"),
			LoadingMsg:  fmt.Sprintf(i18n.T("tag.delete.remote.loading"), selectedTag),
			SuccessMsg:  fmt.Sprintf(i18n.T("tag.delete.remote.success"), selectedTag),
		},
	}

	return command.RunMultipleCommands(commands)
}
