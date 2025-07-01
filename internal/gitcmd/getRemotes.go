package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/logs"
)

func GetRemotes() error {
	output, err := exec.Command("git", "remote", "-v").CombinedOutput()
	if err != nil {
		logs.Error("Failed to get remotes: " + string(output))
		return fmt.Errorf("Failed to get remotes: %s", string(output))
	}
	logs.Success("Remotes:")
	fmt.Println(strings.TrimSpace(string(output)))
	return nil
}
