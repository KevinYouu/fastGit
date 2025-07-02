package command

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/KevinYouu/fastGit/internal/theme"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProgressModel 多步骤进度显示模型
type ProgressModel struct {
	commands     []CommandInfo
	currentStep  int
	total        int
	status       string
	isCompleted  bool
	hasError     bool
	errorMessage string
	results      []string
	executing    bool
}

// StepStartMsg 步骤开始消息
type StepStartMsg struct {
	Step        int
	Description string
}

// StepCompleteMsg 步骤完成消息
type StepCompleteMsg struct {
	Step    int
	Success bool
	Output  string
	Error   error
}

// AllCompleteMsg 所有步骤完成消息
type AllCompleteMsg struct {
	Success bool
}

// NewProgressModel 创建新的进度模型
func NewProgressModel(commands []CommandInfo) *ProgressModel {
	return &ProgressModel{
		commands:    commands,
		currentStep: 0,
		total:       len(commands),
		status:      "Preparing...",
		isCompleted: false,
		results:     make([]string, len(commands)),
		executing:   false,
	}
}

// Init 初始化
func (m *ProgressModel) Init() tea.Cmd {
	return m.executeNextCommand()
}

// Update 更新状态
func (m *ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case StepStartMsg:
		m.currentStep = msg.Step
		m.status = fmt.Sprintf("Executing: %s", msg.Description)
		m.executing = true
		// 开始执行命令
		return m, m.executeCommand(msg.Step)

	case StepCompleteMsg:
		m.results[msg.Step] = msg.Output
		if msg.Success {
			m.status = fmt.Sprintf("Completed: %s", m.commands[msg.Step].Description)
			m.currentStep = msg.Step + 1 // 更新到下一步
			// 继续下一个命令
			if msg.Step+1 < m.total {
				return m, m.executeNextCommand()
			} else {
				// 所有命令完成
				return m, func() tea.Msg { return AllCompleteMsg{Success: true} }
			}
		} else {
			// 命令失败
			m.hasError = true
			m.errorMessage = msg.Error.Error()
			m.status = fmt.Sprintf("Failed: %s", m.commands[msg.Step].Description)
			return m, func() tea.Msg { return AllCompleteMsg{Success: false} }
		}

	case AllCompleteMsg:
		m.isCompleted = true
		m.executing = false
		if msg.Success {
			m.status = "All commands completed successfully!"
		}
		// 等待一秒让用户看到结果，然后退出
		return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
			return tea.Quit()
		})

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	return m, nil
}

// View 渲染视图
func (m *ProgressModel) View() string {
	var s strings.Builder

	// 标题
	titleStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Bold(true).
		Padding(1, 2).
		MarginBottom(1)
	s.WriteString(titleStyle.Render("Executing commands..."))
	s.WriteString("\n\n")

	// 进度条
	progress := float64(m.currentStep) / float64(m.total)
	if m.isCompleted {
		progress = 1.0
	}
	width := 40
	filled := int(progress * float64(width))

	progressBar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			progressBar += "█"
		} else {
			progressBar += "░"
		}
	}

	s.WriteString(fmt.Sprintf("%s [%s] %.0f%% (%d/%d)\n",
		theme.InfoStyle.Render("Progress:"),
		theme.ProgressStyle.Render(progressBar),
		progress*100,
		m.currentStep,
		m.total))

	// 当前状态
	statusIcon := "⏳"
	if m.isCompleted {
		if m.hasError {
			statusIcon = "❌"
		} else {
			statusIcon = "✅"
		}
	} else if m.executing {
		statusIcon = "⚡"
	}

	s.WriteString(fmt.Sprintf("%s %s %s\n",
		theme.InfoStyle.Render("Status:"),
		statusIcon,
		lipgloss.NewStyle().Foreground(theme.TextSecondary).Render(m.status)))

	// 显示当前执行的步骤列表
	s.WriteString("\n")
	for i, cmd := range m.commands {
		var icon string
		var style lipgloss.Style

		if i < m.currentStep {
			icon = "✓"
			style = theme.SuccessStyle
		} else if i == m.currentStep && m.executing {
			icon = "▶"
			style = theme.InfoStyle
		} else if i == m.currentStep && m.hasError {
			icon = "✗"
			style = theme.ErrorStyle
		} else {
			icon = "○"
			style = lipgloss.NewStyle().Foreground(theme.TextSecondary)
		}

		s.WriteString(fmt.Sprintf("  %s %s\n",
			style.Render(icon),
			style.Render(fmt.Sprintf("Step %d: %s", i+1, cmd.Description))))
	}

	// 错误信息
	if m.hasError && m.errorMessage != "" {
		s.WriteString("\n")
		s.WriteString(fmt.Sprintf("%s %s\n",
			theme.ErrorStyle.Render("Error details:"),
			theme.ErrorStyle.Render(m.errorMessage)))
	}

	// 完成时的提示
	if m.isCompleted {
		s.WriteString("\n")
		if m.hasError {
			s.WriteString(theme.ErrorStyle.Render("Process failed. Press any key to exit."))
		} else {
			s.WriteString(theme.SuccessStyle.Render("🎉 All commands completed successfully!"))
		}
	}

	return s.String()
}

// executeNextCommand 执行下一个命令
func (m *ProgressModel) executeNextCommand() tea.Cmd {
	if m.currentStep >= len(m.commands) {
		return func() tea.Msg { return AllCompleteMsg{Success: true} }
	}

	cmd := m.commands[m.currentStep]
	step := m.currentStep

	return func() tea.Msg {
		// 发送开始消息
		return StepStartMsg{
			Step:        step,
			Description: cmd.Description,
		}
	}
}

// executeCommand 执行具体的命令
func (m *ProgressModel) executeCommand(step int) tea.Cmd {
	cmd := m.commands[step]

	return func() tea.Msg {
		// 执行命令
		output, err := exec.Command(cmd.Command, cmd.Args...).CombinedOutput()
		return StepCompleteMsg{
			Step:    step,
			Success: err == nil,
			Output:  string(output),
			Error:   err,
		}
	}
}

// RunMultipleCommandsWithBubbleTea 使用 Bubble Tea 执行多个命令
func RunMultipleCommandsWithBubbleTea(commands []CommandInfo) error {
	model := NewProgressModel(commands)
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed to run progress UI: %w", err)
	}

	// 检查最终状态
	if progressModel, ok := finalModel.(*ProgressModel); ok {
		if progressModel.hasError {
			return fmt.Errorf("%s", progressModel.errorMessage)
		}
	}

	return nil
}
