package gitcmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/pkg/components/colors"
)

type FileStatus struct {
	Status string
	Path   string
}

func statusColor(status string) string {
	switch status {
	case "M":
		return "yellow"
	case "A":
		return "green"
	case "D":
		return "red"
	case "U":
		return "green"
	case "??":
		return "green"
	default:
		return "white"
	}
}

func getFileStatuses() ([]FileStatus, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing command: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var files []FileStatus

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			status := fields[0]
			path := strings.Join(fields[1:], " ")
			files = append(files, FileStatus{Status: status, Path: path})
		}
	}

	return files, nil
}

func Status() {
	fileStatuss, err := getFileStatuses()
	if err != nil {
		fmt.Println(colors.RenderColor("red", "Failed to get file statuses:"), err)
		os.Exit(1)
	}

	if len(fileStatuss) == 0 {
		fmt.Println(colors.RenderColor("blue", "No files changed."))
		os.Exit(0)
	}

	fmt.Println("File statuses:")
	for _, file := range fileStatuss {
		color := statusColor(file.Status)
		fmt.Println(colors.RenderColor(color, file.Status+" "+file.Path))
	}
}
