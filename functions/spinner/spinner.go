package spinner

import (
	"fmt"

	"github.com/charmbracelet/huh/spinner"
)

func Spinner(title, success string, makeBurger func()) error {
	err := spinner.New().
		Title(title).
		Action(makeBurger).
		Run()

	fmt.Println(success)

	return err
}
