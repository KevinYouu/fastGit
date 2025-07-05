package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/i18n"
	"github.com/KevinYouu/fastGit/internal/logs"
)

func GetRemotes() error {
	output, err := exec.Command("git", "remote", "-v").CombinedOutput()
	if err != nil {
		logs.Error(i18n.T("git.remotes.failed") + string(output))
		return fmt.Errorf(i18n.T("git.remotes.failed")+"%s", string(output))
	}
	fmt.Println(i18n.T("git.remotes.title"))
	fmt.Println(strings.TrimSpace(string(output)))
	return nil
}
