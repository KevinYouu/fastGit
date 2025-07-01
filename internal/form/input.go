package form

import (
	"errors"
	"log"

	"github.com/charmbracelet/huh"
)

func Input(title string, defaultValue string) (string, error) {
	inputValue := defaultValue

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title(title).
				Value(&inputValue).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("input cannot be empty")
					}
					return nil
				}),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return inputValue, nil
}
