package remote

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/functions/logs"
)

func GetRemotes() {
	output, err := exec.Command("git", "remote", "-v").CombinedOutput()
	if err != nil {
		logs.Error("Failed to get remotes: " + string(output))
		return
	}
	logs.Success("Remotes:")
	fmt.Println(strings.TrimSpace(string(output)))
}
