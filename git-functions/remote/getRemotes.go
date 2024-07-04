package remote

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
	"github.com/KevinYouu/fastGit/functions/logs"
)

func GetRemotes() {
	output, err := command.RunCmd("git", []string{"remote", "-v"}, "Failed to get remotes: ")
	if err != nil {
		logs.Error("Failed to commit: " + output)
		return
	}

	if output == "" {
		fmt.Println(colors.RenderColor("red", "No remotes found."))
		os.Exit(1)
	}
	fmt.Println(output)
}
