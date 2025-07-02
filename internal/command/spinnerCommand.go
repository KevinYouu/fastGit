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

// RunMultipleCommands 执行多个命令并显示整体进度
func RunMultipleCommands(commands []CommandInfo) error {
	total := len(commands)

	// fmt.Printf("%s\n",
	// 	theme.TitleStyle.Render("Executing commands..."))

	for i, cmd := range commands {
		progress := float64(i) / float64(total)

		// 显示整体进度
		fmt.Printf("\n%s Step %d/%d: %s\n",
			theme.InfoStyle.Render("📋"),
			i+1, total,
			theme.DescriptionStyle.Render(cmd.Description))

		// 显示进度条
		width := 40
		filled := int(progress * float64(width))
		progressBar := ""
		for j := 0; j < width; j++ {
			if j < filled {
				progressBar += "█"
			} else {
				progressBar += "░"
			}
		}

		fmt.Printf("%s %.1f%%\n",
			theme.ProgressStyle.Render(progressBar),
			progress*100)

		// 执行命令
		_, err := RunCmdWithSpinner(cmd.Command, cmd.Args, cmd.LoadingMsg, cmd.SuccessMsg)
		if err != nil {
			return fmt.Errorf("command failed at step %d: %w", i+1, err)
		}
	}

	// 显示最终完成状态
	width := 40
	progressBar := strings.Repeat("█", width)
	fmt.Printf("\n%s 100.0%%\n",
		theme.ProgressStyle.Render(progressBar))

	fmt.Printf("%s %s\n",
		theme.SuccessStyle.Render("🎉"),
		theme.SuccessStyle.Render("All commands completed successfully!"))

	return nil
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

// RunMultiStepCommand 执行多步骤命令
func RunMultiStepCommand(steps []MultiStepInfo) error {
	stepNames := make([]string, len(steps))
	stepMessages := make([]string, len(steps))

	for i, step := range steps {
		stepNames[i] = step.Name
		stepMessages[i] = step.Description
	}

	// 创建多步骤 spinner
	multiSpinner := spinner.NewMultiStepSpinner(stepNames, stepMessages)

	// 创建 tea 程序
	p := tea.NewProgram(multiSpinner.GetSpinner())

	done := make(chan bool)
	errorChan := make(chan error)

	// 在后台执行所有步骤
	go func() {
		var finalErr error

		for i, step := range steps {
			// 更新当前步骤
			multiSpinner.ExecuteStep(i, step.LoadingMsg)

			// 执行命令
			_, err := exec.Command(step.Command, step.Args...).CombinedOutput()
			if err != nil {
				multiSpinner.Complete(false, fmt.Sprintf("Failed at step: %s", step.Name), err)
				finalErr = fmt.Errorf("step %d (%s) failed: %w", i+1, step.Name, err)
				break
			}

			// 添加延迟让用户看到进度
			time.Sleep(500 * time.Millisecond)
		}

		if finalErr == nil {
			multiSpinner.Complete(true, "All steps completed successfully!", nil)
		}

		// 等待一秒让用户看到最终结果
		time.Sleep(2 * time.Second)

		done <- true
		errorChan <- finalErr
	}()

	// 启动 tea 程序显示动画
	go func() {
		p.Run()
	}()

	// 等待完成
	<-done
	p.Quit()

	return <-errorChan
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
