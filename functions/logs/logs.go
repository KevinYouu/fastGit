package logs

import (
	"fmt"

	"github.com/KevinYouu/fastGit/functions/colors"
)

// Waring should be used to render warning text
func Waring(text string) int {
    n, _ := fmt.Println(colors.RenderColor("yellow", text))
    return n
}

func Info(text string) int {
    n, _ := fmt.Println(colors.RenderColor("cyan", text))
    return n
}

func Success(text string) int {
    n, _ := fmt.Println(colors.RenderColor("green", text))
    return n
}

func Error(text string) int {
    n, _ := fmt.Println(colors.RenderColor("red", text))
    return n
}