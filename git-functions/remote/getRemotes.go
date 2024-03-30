package remote

import (
	"fmt"
	"os"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/KevinYouu/fastGit/functions/command"
)

func GetRemotes() {
	addlog, err := command.RunCommand("git", "remote", "-v")
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get remotes: "+err.Error()))
		os.Exit(1)
	}

	if addlog == "" {
		fmt.Println(colors.RenderColor("red", "No remotes found."))
		os.Exit(1)
	}
	fmt.Println(addlog)
}
