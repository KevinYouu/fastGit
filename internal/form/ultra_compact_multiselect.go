package form

import (
	"os"
	"strings"

	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// UltraCompactMultiSelectForm 超紧凑多选表单
func UltraCompactMultiSelectForm(title string, options []string) ([]string, error) {
	var selectedValues []string
	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	// 获取终端宽度和高度
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 24
	}

	// 计算多选区域的高度
	maxHeight := height - 2 // 预留标题和底部空间
	if maxHeight < 3 {
		maxHeight = 3
	}
	if maxHeight > 8 {
		maxHeight = 8
	}

	selectHeight := min(len(options), maxHeight)

	// 创建带有右侧提示信息的标题
	helpText := "↑/↓:导航 Space:选择 Enter:确认"
	titleWithHelp := createMultiSelectTitleWithRightHelp(title, helpText, width)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Height(selectHeight).
				Title(titleWithHelp).
				Description(""). // 清空描述
				Options(selectOpts...).
				Value(&selectedValues),
		),
	).WithTheme(theme.GetUltraCompactTheme())

	err = form.Run()
	if err != nil {
		return nil, err
	}

	return selectedValues, nil
}

// createMultiSelectTitleWithRightHelp 为多选组件创建带有右侧帮助信息的标题
func createMultiSelectTitleWithRightHelp(title, helpText string, terminalWidth int) string {
	if terminalWidth < 60 {
		// 终端太窄，只显示标题
		return title
	}

	titleLen := len(title)
	helpLen := len(helpText)

	// 计算需要的空格数
	spacesNeeded := terminalWidth - titleLen - helpLen - 4 // 预留4个字符的缓冲
	if spacesNeeded < 2 {
		spacesNeeded = 2
	}

	// 创建标题和帮助信息的布局
	titleStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Bold(true)

	helpStyle := lipgloss.NewStyle().
		Foreground(theme.TextSecondary).
		Bold(false)

	// 使用 lipgloss 创建左右对齐的布局
	leftPart := titleStyle.Render(title)
	rightPart := helpStyle.Render(helpText)

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		leftPart,
		strings.Repeat(" ", spacesNeeded),
		rightPart,
	)
}
