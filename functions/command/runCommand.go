package command

import (
	"fmt"
	"os/exec"
)

// RunCommand 执行给定的命令和参数，并输出执行结果
func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v", err)
	}

	return string(output), nil
}
