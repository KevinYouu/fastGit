package form

import (
	"github.com/charmbracelet/huh"
)

type Option struct {
	Label string
	Value string
}

func SelectForm(title string, options []Option) (label, value string, err error) {
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt.Label, opt.Value)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(100).
				Title(title).
				Options(selectOpts...).
				Value(&selectedValue),
		))
	err = form.Run()
	if err != nil {
		return "", "", err
	}

	// Find the label for the selected value
	for _, opt := range options {
		if opt.Value == selectedValue {
			label = opt.Label
			break
		}
	}

	return label, selectedValue, nil
}
