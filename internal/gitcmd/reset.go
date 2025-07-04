package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

type Commit struct {
	Hash    string
	Message string
	Date    string
	Author  string
	Email   string
	IsHead  bool
}

func Reset() error {
	// 显示开始信息 - 简洁的标题
	headerStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Bold(true).
		Padding(0, 1)

	fmt.Printf("%s\n", headerStyle.Render("🔄 Git Reset"))

	// 使用更详细的git log格式获取提交历史
	cmd := exec.Command("git", "log", "--pretty=format:%h|%s|%ad|%an|%ae", "--date=format:%m-%d %H:%M")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error executing git log command: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var options = []config.Option{}
	var commits = []Commit{}

	// 解析并存储提交信息（不显示历史记录）
	for i, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) == 5 {
			hash := parts[0]
			message := parts[1]
			date := parts[2]
			author := parts[3]
			email := parts[4]

			// 存储提交信息
			commits = append(commits, Commit{
				Hash:    hash,
				Message: message,
				Date:    date,
				Author:  author,
				Email:   email,
				IsHead:  i == 0,
			})

			// 限制消息长度，避免过长
			shortMsg := message
			if len(shortMsg) > 40 {
				shortMsg = shortMsg[:37] + "..."
			}

			// 添加到选择列表，使用纯文本格式以允许背景色正确覆盖
			commitLabel := ""
			if i == 0 {
				// HEAD提交使用纯文本格式，但添加标记以区分
				commitLabel = fmt.Sprintf(
					"[HEAD] %s %s\n%s • %s",
					hash,
					shortMsg,
					date,
					author,
				)
			} else {
				// 普通提交使用纯文本格式
				commitLabel = fmt.Sprintf(
					"%s %s\n%s • %s",
					hash,
					shortMsg,
					date,
					author,
				)
			}
			options = append(options, config.Option{Label: commitLabel, Value: hash})
		}
	}

	// 使用表格选择表单
	_, choose, err := form.TableSelectForm(options)
	if err != nil {
		return fmt.Errorf("选择提交错误: %w", err)
	}

	// 获取选择的提交完整信息
	var selectedCommit Commit
	for _, commit := range commits {
		if commit.Hash == choose {
			selectedCommit = commit
			break
		}
	}

	// 选择重置模式，使用更紧凑的格式 - 纯文本格式以确保背景色能正确覆盖
	resetModes := []config.Option{
		{
			Label: "Soft - 保留工作目录和暂存区",
			Value: "--soft",
		},
		{
			Label: "Mixed - 保留工作目录，清空暂存区",
			Value: "--mixed",
		},
		{
			Label: "Hard - 丢弃所有未提交的更改",
			Value: "--hard",
		},
	}

	// 使用表格选择表单选择重置模式
	_, resetMode, err := form.TableSelectForm(resetModes)
	if err != nil {
		return fmt.Errorf("选择重置模式错误: %w", err)
	}

	// 获取可读的重置模式名称
	resetModeReadable := strings.TrimPrefix(resetMode, "--")

	// 根据重置模式选择对应的颜色
	var modeColor lipgloss.Style
	switch resetMode {
	case "--soft":
		modeColor = lipgloss.NewStyle().Foreground(theme.InfoColor)
	case "--mixed":
		modeColor = lipgloss.NewStyle().Foreground(theme.WarningColor)
	case "--hard":
		modeColor = lipgloss.NewStyle().Foreground(theme.ErrorColor)
	}

	// 构建更紧凑的确认信息
	shortMsg := selectedCommit.Message
	if len(shortMsg) > 40 {
		shortMsg = shortMsg[:37] + "..."
	}

	confirmDesc := fmt.Sprintf("确认重置到 %s  "+"%s "+
		"%s模式 %s",
		lipgloss.NewStyle().Foreground(theme.PrimaryColor).Bold(true).Render(selectedCommit.Hash),
		shortMsg,
		modeColor.Render(resetModeReadable),
		getModeDescription(resetMode),
	)

	// 针对 hard 模式添加警告，但更紧凑
	if resetMode == "--hard" {
		confirmDesc += "\n" + lipgloss.NewStyle().
			Foreground(theme.ErrorColor).
			Bold(true).
			Render("⚠️ 将丢失所有未提交更改！")
	}

	// 使用自定义确认表单
	confirm := form.Confirm(confirmDesc)

	if confirm {
		// 执行重置操作
		resetArgs := []string{"reset"}
		if resetMode != "--mixed" { // mixed是默认值，不需要显式指定
			resetArgs = append(resetArgs, resetMode)
		}
		resetArgs = append(resetArgs, choose)

		_, err := command.RunCmdWithSpinner("git", resetArgs,
			fmt.Sprintf("正在重置 (%s)...", resetModeReadable),
			fmt.Sprintf("已重置到 %s (%s)", choose, resetModeReadable))
		if err != nil {
			return fmt.Errorf("执行git reset命令时出错: %w", err)
		}

		// 显示简洁的成功信息
		fmt.Printf("\n%s %s\n",
			theme.SuccessStyle.Render("✓"),
			lipgloss.NewStyle().
				Foreground(theme.SuccessColor).
				Render(fmt.Sprintf("重置完成 (HEAD → %s)", choose)))

		// 简洁的操作提示
		switch resetMode {
		case "--soft":
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.InfoColor).
					Render("💡 更改已保留在暂存区"))
		case "--mixed":
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.InfoColor).
					Render("💡 更改已保留在工作区"))
		case "--hard":
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.InfoColor).
					Render("💡 所有未提交更改已丢弃"))
		}
	} else {
		fmt.Printf("\n%s %s\n",
			theme.InfoStyle.Render("ℹ️"),
			theme.InfoStyle.Render("已取消"))
	}
	return nil
}

// 获取重置模式的简短描述
func getModeDescription(mode string) string {
	switch mode {
	case "--soft":
		return lipgloss.NewStyle().
			Foreground(theme.TextSecondary).
			Render(" (保留全部)")
	case "--mixed":
		return lipgloss.NewStyle().
			Foreground(theme.TextSecondary).
			Render(" (默认)")
	case "--hard":
		return lipgloss.NewStyle().
			Foreground(theme.TextSecondary).
			Render(" (危险)")
	default:
		return ""
	}
}
