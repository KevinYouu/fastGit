package status

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"
)

// FileStatus 结构体表示文件的状态和路径
type FileStatus struct {
	Status string // 文件状态 (如 "M", "A", "??")
	Path   string // 文件路径
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

func GetFileStatuses() ([]FileStatus, error) {
	// 执行 git status 命令
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing command: %v", err)
	}

	// 解析 git status 输出并创建 FileStatus 结构体实例
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

	fmt.Println("File statuses:")
	for _, file := range files {
		color := statusColor(file.Status)
		fmt.Println(colors.RenderColor(color, file.Status+" "+file.Path))
	}

	return files, nil
}
