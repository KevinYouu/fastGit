package form

import (
	"os"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
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

	// 获取终端高度来调整布局
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24
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

	// 创建标题（不包含帮助信息）
	titleWithHelp := title

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(selectHeight).
				Title(titleWithHelp).
				Description(""). // 清空描述，因为已经在标题中显示
				Options(selectOpts...).
				Value(&selectedValue).
				Filtering(false), // 禁用过滤功能和相关提示
		),
	).WithTheme(theme.GetUltraCompactTheme()).
		WithShowHelp(false) // 禁用帮助信息

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

	// 获取终端高度来调整布局
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24
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

	// 创建标题（不包含帮助信息）
	titleWithHelp := title

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(selectHeight).
				Title(titleWithHelp).
				Description(""). // 清空描述
				Options(selectOpts...).
				Value(&selectedValue).
				Filtering(false), // 禁用过滤功能和相关提示
		),
	).WithTheme(theme.GetUltraCompactTheme()).
		WithShowHelp(false) // 禁用帮助信息

	err = form.Run()
	if err != nil {
		return "", "", err
	}

	return selectedValue, selectedValue, nil
}
