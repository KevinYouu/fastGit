package form

import (
	"github.com/KevinYouu/fastGit/functions/config"
	"github.com/charmbracelet/huh"
)

func SelectForm(title string, options []config.Option) (label, value string, err error) {
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

func SelectFormWithStringSlice(title string, options []string) (label, value string, err error) {
	var selectedValue string

	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
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
		if opt == selectedValue {
			label = opt
			break
		}
	}

	return label, selectedValue, nil
}
