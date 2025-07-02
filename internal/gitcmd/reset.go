package gitcmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/internal/command"
	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/form"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/lipgloss"
)

type Commit struct {
	Hash    string
	Message string
	Date    string
	Author  string
	Email   string
}

func Reset() error {
	// 显示开始信息
	fmt.Printf("%s %s\n",
		theme.InfoStyle.Render("🔄"),
		theme.TitleStyle.Render("Git Reset to Previous Commit"))

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
				"%s %s %s %s <%s>",
				lipgloss.NewStyle().Foreground(theme.PrimaryColor).Render(hash),
				lipgloss.NewStyle().Foreground(theme.TextColor).Render(message),
				lipgloss.NewStyle().Foreground(theme.TextSecondary).Render(date),
				lipgloss.NewStyle().Foreground(theme.AccentColor).Render(author),
				lipgloss.NewStyle().Foreground(theme.TextMuted).Render(email),
			), Value: hash})
		}
	}

	_, choose, err := form.SelectForm("Choose a commit to reset to", options)
	if err != nil {
		return fmt.Errorf("SelectForm: %w", err)
	}

	confirm := form.Confirm(fmt.Sprintf("Are you sure you want to reset to commit %s?", choose))
	if confirm {
		_, err := command.RunCmdWithSpinner("git", []string{"reset", choose},
			fmt.Sprintf("Resetting to commit %s...", choose),
			fmt.Sprintf("Successfully reset to commit %s", choose))
		if err != nil {
			return fmt.Errorf("Error executing git reset command: %w", err)
		}
	} else {
		fmt.Printf("%s %s\n",
			theme.InfoStyle.Render("ℹ️"),
			theme.InfoStyle.Render("Reset operation cancelled"))
	}
	return nil
}
