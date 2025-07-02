package form

import (
	"errors"

	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/huh"
)

func Input(title string, defaultValue string) (string, error) {
	inputValue := defaultValue

	// 直接使用紧凑模式
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(title).
				Description("输入后按 Enter").
				Placeholder("请输入...").
				Value(&inputValue).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("输入不能为空")
					}
					return nil
				}),
		),
	).WithTheme(theme.GetCompactTheme(true))

	err := form.Run()
	if err != nil {
		return "", err
	}

	return inputValue, nil
}
