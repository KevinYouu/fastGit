package gitcmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/pkg/components/colors"
	"github.com/KevinYouu/fastGit/pkg/components/config"
	"github.com/KevinYouu/fastGit/pkg/components/form"
)

// Commit struct represents a commit record
type Commit struct {
	Hash    string // Commit hash
	Message string // Commit message
	Date    string // Commit date
	Author  string // Commit author
	Email   string // Commit author email
}

func Reset() {
	// Execute git log command
	cmd := exec.Command("git", "log", "--pretty=format:%h|%s|%ad|%an|%ae", "--date=short")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing git log command:", err)
		return
	}

	// Parse git log output and create Commit struct instances
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
		fmt.Println(err)
		os.Exit(1)
	}
	confirm := form.Confirm("Reset to commit: " + choose + " ?")
	if confirm {
		cmd = exec.Command("git", "reset", choose)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println("Error executing git log command:", err)
			return
		}
		fmt.Println(colors.RenderColor("blue", "Reset to commit: "+choose))
	}
}
