package form

import (
	"os"
	"strings"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// UltraCompactSelectForm 超紧凑选择表单，适用于极低高度终端
// 将提示信息显示在右侧，最大化纵向空间
func UltraCompactSelectForm(title string, options []config.Option) (label, value string, err error) {
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt.Label, opt.Value)
	}

	// 获取终端宽度来调整布局
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 24
	}

	// 计算选择区域的高度
	maxHeight := height - 2 // 预留标题和底部空间
	if maxHeight < 3 {
		maxHeight = 3
	}
	if maxHeight > 10 {
		maxHeight = 10
	}

	selectHeight := min(len(options), maxHeight)

	// 创建带有右侧提示信息的标题
	helpText := "↑/↓:选择 Enter:确认 q:退出"
	titleWithHelp := createTitleWithRightHelp(title, helpText, width)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(selectHeight).
				Title(titleWithHelp).
				Description(""). // 清空描述，因为已经在标题中显示
				Options(selectOpts...).
				Value(&selectedValue),
		),
	).WithTheme(theme.GetUltraCompactTheme())

	err = form.Run()
	if err != nil {
		return "", "", err
	}

	// 找到选中的选项
	for _, opt := range options {
		if opt.Value == selectedValue {
			return opt.Label, opt.Value, nil
		}
	}

	return "", "", nil
}

// UltraCompactSelectFormWithStringSlice 超紧凑字符串选择表单
func UltraCompactSelectFormWithStringSlice(title string, options []string) (label, value string, err error) {
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	// 获取终端宽度来调整布局
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width, height = 80, 24
	}

	// 计算选择区域的高度
	maxHeight := height - 2 // 预留标题和底部空间
	if maxHeight < 3 {
		maxHeight = 3
	}
	if maxHeight > 10 {
		maxHeight = 10
	}

	selectHeight := min(len(options), maxHeight)

	// 创建带有右侧提示信息的标题
	helpText := "↑/↓:选择 Enter:确认 q:退出"
	titleWithHelp := createTitleWithRightHelp(title, helpText, width)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(selectHeight).
				Title(titleWithHelp).
				Description(""). // 清空描述
				Options(selectOpts...).
				Value(&selectedValue),
		),
	).WithTheme(theme.GetUltraCompactTheme())

	err = form.Run()
	if err != nil {
		return "", "", err
	}

	return selectedValue, selectedValue, nil
}

// createTitleWithRightHelp 创建带有右侧帮助信息的标题
func createTitleWithRightHelp(title, helpText string, terminalWidth int) string {
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
