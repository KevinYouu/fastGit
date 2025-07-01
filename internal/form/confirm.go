package form

import (
	"log"

	"github.com/charmbracelet/huh"
)

func Confirm(title string) bool {
	var discount bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(title).
				Value(&discount),
		),
	)
	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return discount
}
