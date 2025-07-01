package form

import (
	"github.com/charmbracelet/huh"
)

func MultiSelectForm(title string, options []string) (Values []string, err error) {
	var selectedValues []string
	selectOpts := make([]huh.Option[string], len(options))
	for i, opt := range options {
		selectOpts[i] = huh.NewOption(opt, opt)
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(title).
				Options(selectOpts...).
				Value(&selectedValues),
		),
	)
	err = form.Run()
	if err != nil {
		return nil, err
	}

	return selectedValues, nil
}
