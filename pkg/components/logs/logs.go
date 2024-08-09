package logs

import (
	"fmt"

	"github.com/KevinYouu/fastGit/pkg/components/colors"
)

// Waring should be used to render warning text
func Waring(text string) {
	fmt.Println(colors.RenderColor("yellow", text))
}

func Info(text string) {
	fmt.Println(colors.RenderColor("cyan", text))
}

func Success(text string) {
	fmt.Println(colors.RenderColor("green", text))
}

func Error(text string) {
	fmt.Println(colors.RenderColor("red", text))
}
