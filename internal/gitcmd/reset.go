package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/colors"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
)

type Commit struct {
	Hash    string
	Message string
	Date    string
	Author  string
	Email   string
}

func Reset() error {
	cmd := exec.Command("git", "log", "--pretty=format:%h|%s|%ad|%an|%ae", "--date=short")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error executing git log command: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var options = []config.Option{}

	for _, line := range lines {
		parts := strings.Split(line, "|")
		if len(parts) == 5 {
			hash := parts[0]
			message := parts[1]
			date := parts[2]
			author := parts[3]
			email := parts[4]
			options = append(options, config.Option{Label: fmt.Sprintf(
				"Hash: %s | Message: %s | Date: %s | Author: %s | Email: %s ",
				hash, message, date, author, email,
			), Value: hash})
		}
	}

	_, choose, err := form.SelectForm("Choose a commit type", options)
	if err != nil {
		return fmt.Errorf("SelectForm: %w", err)
	}
	confirm := form.Confirm("Reset to commit: " + choose + " ?")
	if confirm {
		cmd = exec.Command("git", "reset", choose)
		_, err = cmd.Output()
		if err != nil {
			return fmt.Errorf("Error executing git reset command: %w", err)
		}
		fmt.Println(colors.RenderColor("blue", "Reset to commit: "+choose))
	}
	return nil
}
