package form

import (
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
)

func Confirm(title string) bool {
	var confirmed bool

	// 直接使用紧凑模式
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Description("←/→ 或 y/n，Enter 确认").
				Value(&confirmed),
		),
	).WithTheme(theme.GetCompactTheme(true))

	err := form.Run()
	if err != nil {
		return false
	}

	return confirmed
}
