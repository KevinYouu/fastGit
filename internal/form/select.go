package form

import (
	"os"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SelectForm(title string, options []config.Option) (label, value string, err error) {
	// 使用统一的紧凑布局
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt.Label, opt.Value)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(title).
				Options(selectOpts...).
				Value(&selectedValue).
				Filtering(false),
		),
	).WithTheme(theme.GetCompactTheme()).
		WithShowHelp(false)

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

func SelectFormWithStringSlice(title string, options []string) (label, value string, err error) {
	// 检测终端高度用于高度计算
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // 默认值
	}

	// 使用统一的紧凑布局
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	// 计算合适的高度
	availableHeight := height - 6
	compactHeight := max(min(len(options)+1, availableHeight), 3)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(compactHeight).
				Title(title).
				Options(selectOpts...).
				Value(&selectedValue).
				Filtering(false),
		),
	).WithTheme(theme.GetCompactTheme()).
		WithShowHelp(false)

	err = form.Run()
	if err != nil {
		return "", "", err
	}

	return selectedValue, selectedValue, nil
}
