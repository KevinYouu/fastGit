package form

import (
	"os"

	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
	"golang.org/x/term"
)

func MultiSelectForm(title string, options []string) (Values []string, err error) {
	// 检测终端高度，决定使用哪种布局
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		height = 24 // 默认值
	}

	// 如果终端高度非常低（< 12行），使用超紧凑布局
	if height < 12 {
		return UltraCompactMultiSelectForm(title, options)
	}

	// 否则使用标准紧凑布局
	var selectedValues []string
	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	// 计算合适的高度，确保多选项可见
	availableHeight := height - 6                         // 预留标题、描述、边框等空间
	compactHeight := min(len(options)+1, availableHeight) // +1 确保有额外空间
	if compactHeight < 3 {
		compactHeight = 3
	}
	if compactHeight > 8 {
		compactHeight = 8 // 限制最大高度
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(title).
				Description("↑/↓ 导航，Space 选择，Enter 确认").
				Height(compactHeight).
				Options(selectOpts...).
				Value(&selectedValues),
		),
	).WithTheme(theme.GetCompactTheme())

	err = form.Run()
	if err != nil {
		return nil, err
	}

	return selectedValues, nil
}
