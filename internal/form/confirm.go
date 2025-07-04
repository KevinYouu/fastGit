package form

import (
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
)

func Confirm(title string) bool {
	var confirmed bool

	// 直接使用紧凑模式，并尝试禁用帮助
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Value(&confirmed),
		),
	).WithTheme(theme.GetCompactTheme()).
		WithShowHelp(false) // 尝试禁用帮助

	err := form.Run()
	if err != nil {
		return false
	}

	return confirmed
}
