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

func SelectForm(title string, options []config.Option) (label, value string, err error) {
	// 检测终端高度，决定使用哪种布局
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // 默认值
	}

	// 如果终端高度非常低（< 12行），使用超紧凑布局
	if height < 12 {
		return UltraCompactSelectForm(title, options)
	}

	// 否则使用标准紧凑布局
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt.Label, opt.Value)
	}

	// 计算合适的高度，确保最后一项可见
	availableHeight := height - 6                         // 预留标题、描述、边框等空间
	compactHeight := min(len(options)+1, availableHeight) // +1 确保有额外空间
	if compactHeight < 3 {
		compactHeight = 3
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(compactHeight).
				Title(title).
				Description("↑/↓ 选择, Enter 确认").
				Options(selectOpts...).
				Value(&selectedValue),
		),
	).WithTheme(theme.GetCompactTheme(true))

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
	// 检测终端高度，决定使用哪种布局
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // 默认值
	}

	// 如果终端高度非常低（< 12行），使用超紧凑布局
	if height < 12 {
		return UltraCompactSelectFormWithStringSlice(title, options)
	}

	// 否则使用标准紧凑布局
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	// 计算合适的高度，确保最后一项可见
	availableHeight := height - 6                         // 预留标题、描述、边框等空间
	compactHeight := min(len(options)+1, availableHeight) // +1 确保有额外空间
	if compactHeight < 3 {
		compactHeight = 3
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(compactHeight).
				Title(title).
				Description("↑/↓ 选择, Enter 确认, q 退出").
				Options(selectOpts...).
				Value(&selectedValue),
		),
	).WithTheme(theme.GetCompactTheme(true))

	err = form.Run()
	if err != nil {
		return "", "", err
	}

	return selectedValue, selectedValue, nil
}
