package command

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/KevinYouu/fastGit/internal/spinner"
	"github.com/KevinYouu/fastGit/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// RunCmdWithSpinner 带加载动画的命令执行
func RunCmdWithSpinner(command string, args []string, loadingMsg, successMsg string) (string, error) {
	// 创建加载动画的channel
	done := make(chan bool)
	result := make(chan string, 1) // 添加缓冲避免阻塞
	errChan := make(chan error, 1) // 添加缓冲避免阻塞

	// 启动加载动画
	go func() {
		frames := theme.GetSpinnerFrames()
		style := theme.GetSpinnerStyle()
		frameIndex := 0
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("\r%s %s",
					style.Render(frames[frameIndex]),
					theme.InfoStyle.Render(loadingMsg))
				frameIndex = (frameIndex + 1) % len(frames)
			}
		}
	}()

	// 在goroutine中执行命令
	go func() {
		defer func() {
			done <- true // 确保动画停止
		}()

		output, err := exec.Command(command, args...).CombinedOutput()

		// 先发送结果，再停止动画
		errChan <- err
		result <- string(output)
	}()

	// 等待命令完成
	err := <-errChan
	output := <-result

	// 清除加载动画行
	fmt.Print("\r" + strings.Repeat(" ", len(loadingMsg)+10) + "\r")

	if err != nil {
		fmt.Printf("%s %s\n",
			theme.ErrorStyle.Render("✗"),
			theme.ErrorStyle.Render(fmt.Sprintf("Failed: %s", loadingMsg)))

		// 显示详细的错误输出
		trimmedOutput := strings.TrimSpace(output)
		if trimmedOutput != "" {
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.ErrorColor).
					Bold(true).
					Render("Error details:"))
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.TextSecondary).
					Render(trimmedOutput))
		}

		return output, err
	}

	// 显示成功信息
	fmt.Printf("%s %s\n",
		theme.SuccessStyle.Render("✓"),
		theme.SuccessStyle.Render(successMsg))

	// 如果有输出内容，显示它
	trimmedOutput := strings.TrimSpace(output)
	if trimmedOutput != "" {
		fmt.Println(lipgloss.NewStyle().Foreground(theme.TextSecondary).Render(trimmedOutput))
	}

	return output, nil
}

// RunCmdWithProgress 带进度显示的命令执行
func RunCmdWithProgress(command string, args []string, loadingMsg, successMsg string, estimatedDuration time.Duration) (string, error) {
	done := make(chan bool)
	result := make(chan string, 1) // 添加缓冲避免阻塞
	errChan := make(chan error, 1) // 添加缓冲避免阻塞

	// 启动进度显示
	go func() {
		frames := theme.GetSpinnerFrames()
		style := theme.GetSpinnerStyle()
		frameIndex := 0
		ticker := time.NewTicker(100 * time.Millisecond)
		defer ticker.Stop()

		startTime := time.Now()
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				elapsed := time.Since(startTime)
				progress := float64(elapsed) / float64(estimatedDuration)
				if progress > 1.0 {
					progress = 1.0
				}

				// 创建进度条
				width := 20
				filled := int(progress * float64(width))
				progressBar := ""
				for i := 0; i < width; i++ {
					if i < filled {
						progressBar += "█"
					} else {
						progressBar += "░"
					}
				}

				fmt.Printf("\r%s %s [%s] %.1f%%",
					style.Render(frames[frameIndex]),
					theme.InfoStyle.Render(loadingMsg),
					lipgloss.NewStyle().Foreground(theme.PrimaryColor).Render(progressBar),
					progress*100)

				frameIndex = (frameIndex + 1) % len(frames)
			}
		}
	}()

	// 在goroutine中执行命令
	go func() {
		defer func() {
			done <- true // 确保动画停止
		}()

		output, err := exec.Command(command, args...).CombinedOutput()

		// 先发送结果，再停止动画
		errChan <- err
		result <- string(output)
	}()

	// 等待命令完成
	err := <-errChan
	output := <-result

	// 清除进度行
	fmt.Print("\r" + strings.Repeat(" ", len(loadingMsg)+50) + "\r")

	if err != nil {
		fmt.Printf("%s %s\n",
			theme.ErrorStyle.Render("✗"),
			theme.ErrorStyle.Render(fmt.Sprintf("Failed: %s", loadingMsg)))

		// 显示详细的错误输出
		trimmedOutput := strings.TrimSpace(output)
		if trimmedOutput != "" {
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.ErrorColor).
					Bold(true).
					Render("Error details:"))
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.TextSecondary).
					Render(trimmedOutput))
		}

		return output, err
	}

	// 显示成功信息
	fmt.Printf("%s %s\n",
		theme.SuccessStyle.Render("✓"),
		theme.SuccessStyle.Render(successMsg))

	// 如果有输出内容，显示它
	trimmedOutput := strings.TrimSpace(output)
	if trimmedOutput != "" {
		fmt.Println(lipgloss.NewStyle().Foreground(theme.TextSecondary).Render(trimmedOutput))
	}

	return output, nil
}

// CommandInfo 命令信息结构
type CommandInfo struct {
	Command     string
	Args        []string
	Description string
	LoadingMsg  string
	SuccessMsg  string
}

// RunCmdWithAdvancedSpinner 使用高级加载动画执行命令
func RunCmdWithAdvancedSpinner(command string, args []string, loadingMsg, successMsg string, spinnerType spinner.AdvancedSpinnerType) (string, error) {
	// 创建高级 spinner
	advSpinner := spinner.NewAdvancedSpinner(spinnerType, loadingMsg)

	// 创建 tea 程序
	p := tea.NewProgram(advSpinner)

	done := make(chan bool)
	result := make(chan string)
	errChan := make(chan error)

	// 在后台执行命令
	go func() {
		output, err := exec.Command(command, args...).CombinedOutput()

		// 更新 spinner 状态
		if err != nil {
			advSpinner.SetDone(false, "Command failed", err)
		} else {
			advSpinner.SetDone(true, successMsg, nil)
		}

		// 等待一秒让用户看到结果
		time.Sleep(time.Second)

		done <- true
		errChan <- err
		result <- string(output)
	}()

	// 启动 tea 程序显示动画
	go func() {
		p.Run()
	}()

	// 等待命令完成
	<-done
	p.Quit()

	output := <-result
	err := <-errChan

	return output, err
}

// RunMultiStepCommand 执行多步骤命令，使用统一的进度条组件
func RunMultiStepCommand(steps []MultiStepInfo) error {
	// 将 MultiStepInfo 转换为 CommandInfo
	commands := make([]CommandInfo, len(steps))
	for i, step := range steps {
		commands[i] = CommandInfo{
			Command:     step.Command,
			Args:        step.Args,
			Description: step.Description,
			LoadingMsg:  step.LoadingMsg,
			SuccessMsg:  fmt.Sprintf("Completed: %s", step.Description),
		}
	}

	// 使用统一的进度条组件
	return RunMultipleCommandsWithBubbleTea(commands)
}

// MultiStepInfo 多步骤信息结构
type MultiStepInfo struct {
	Name        string
	Description string
	Command     string
	Args        []string
	LoadingMsg  string
}

// RunCmdWithCustomSpinner 使用自定义 spinner 样式执行命令
func RunCmdWithCustomSpinner(command string, args []string, message string, spinnerFrames []string) (string, error) {
	done := make(chan bool)
	result := make(chan string, 1) // 添加缓冲避免阻塞
	errChan := make(chan error, 1) // 添加缓冲避免阻塞

	// 启动自定义动画
	go func() {
		style := theme.GetSpinnerStyle()
		frameIndex := 0
		ticker := time.NewTicker(150 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("\r%s %s",
					style.Render(spinnerFrames[frameIndex]),
					theme.InfoStyle.Render(message))
				frameIndex = (frameIndex + 1) % len(spinnerFrames)
			}
		}
	}()

	// 执行命令
	go func() {
		defer func() {
			done <- true // 确保动画停止
		}()

		output, err := exec.Command(command, args...).CombinedOutput()

		// 先发送结果，再停止动画
		errChan <- err
		result <- string(output)
	}()

	// 等待完成
	err := <-errChan
	output := <-result

	// 清除动画行
	fmt.Print("\r" + strings.Repeat(" ", len(message)+10) + "\r")

	if err != nil {
		fmt.Printf("%s %s\n",
			theme.ErrorStyle.Render(theme.GetStatusIcon("error")),
			theme.ErrorStyle.Render(fmt.Sprintf("Failed: %s", message)))

		// 显示详细的错误输出
		trimmedOutput := strings.TrimSpace(output)
		if trimmedOutput != "" {
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.ErrorColor).
					Bold(true).
					Render("Error details:"))
			fmt.Printf("%s\n",
				lipgloss.NewStyle().
					Foreground(theme.TextSecondary).
					Render(trimmedOutput))
		}

		return output, err
	}

	fmt.Printf("%s %s\n",
		theme.SuccessStyle.Render(theme.GetStatusIcon("success")),
		theme.SuccessStyle.Render(message))

	return output, nil
}
