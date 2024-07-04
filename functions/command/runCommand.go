package command

import (
	"fmt"
	"os/exec"
)

// RunCommand 执行给定的命令和参数，并输出执行结果
func RunCommand(command string, args ...string) (string, error) {
	output, err := exec.Command(command, args...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v", err)
	}

	return string(output), nil
}
