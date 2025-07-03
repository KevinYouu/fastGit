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

// CommandStepModel 表示使用 Bubble Tea 的命令步骤执行模型
type CommandStepModel struct {
	// 命令列表
	Commands []CommandInfo

	// 当前状态
	currentStep int
	totalSteps  int
	spinner     spinnerAnimation
	frame       int
	done        bool
	err         error
	output      []string
	stepStatus  []int // 0=pending, 1=running, 2=success, 3=failed
}

// spinnerAnimation 实现简单的加载动画
type spinnerAnimation struct {
	frames []string
	fps    time.Duration
}

// 默认加载动画
var defaultSpinnerAnimation = spinnerAnimation{
	frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	fps:    time.Second / 10,
}

// 初始化模型
func NewCommandStepModel(commands []CommandInfo) *CommandStepModel {
	return &CommandStepModel{
		Commands:    commands,
		currentStep: 0,
		totalSteps:  len(commands),
		spinner:     defaultSpinnerAnimation,
		output:      make([]string, len(commands)),
		stepStatus:  make([]int, len(commands)),
	}
}

// 初始化开始执行
func (m *CommandStepModel) Init() tea.Cmd {
	return tea.Batch(
		m.tickCmd(),
		m.runCommand(),
	)
}

// 帧更新命令
func (m *CommandStepModel) tickCmd() tea.Cmd {
	return tea.Tick(m.spinner.fps, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// 消息类型
type tickMsg time.Time
type stepCompleteMsg struct {
	index  int
	err    error
	output string
}

// 更新模型状态
func (m *CommandStepModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil

	case tickMsg:
		// 更新加载动画帧
		m.frame = (m.frame + 1) % len(m.spinner.frames)
		if m.done {
			return m, tea.Tick(time.Second, func(time.Time) tea.Msg {
				return tea.Quit()
			})
		}
		return m, m.tickCmd()

	case stepCompleteMsg:
		// 处理步骤完成
		m.output[msg.index] = msg.output

		if msg.err != nil {
			// 步骤失败
			m.stepStatus[msg.index] = 3 // 失败
			m.err = msg.err
			m.done = true
			return m, nil
		}

		// 步骤成功
		m.stepStatus[msg.index] = 2 // 成功

		// 移动到下一步
		next := msg.index + 1
		if next < m.totalSteps {
			m.currentStep = next
			m.stepStatus[next] = 1 // 标记为正在运行
			return m, m.runCommand()
		}

		// 全部完成
		m.done = true
		return m, nil
	}

	return m, nil
}

// 渲染视图
func (m *CommandStepModel) View() string {
	var s strings.Builder

	// 标题
	s.WriteString(lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.PrimaryColor).
		Render("🚀 Executing Commands"))

	// 进度条
	progress := float64(m.currentStep) / float64(m.totalSteps)
	if m.done && m.err == nil {
		progress = 1.0
	}

	width := 40
	filled := int(progress * float64(width))
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	barStyle := lipgloss.NewStyle().Foreground(theme.PrimaryColor)
	s.WriteString(fmt.Sprintf("Progress: [%s] %.0f%% (%d/%d)\n\n",
		barStyle.Render(bar),
		progress*100,
		m.currentStep,
		m.totalSteps))

	// 步骤列表
	for i, cmd := range m.Commands {
		var icon string
		var style lipgloss.Style

		switch m.stepStatus[i] {
		case 0: // 等待
			icon = "○"
			style = lipgloss.NewStyle().Foreground(theme.TextSecondary)
		case 1: // 运行中
			icon = m.spinner.frames[m.frame]
			style = lipgloss.NewStyle().Foreground(theme.PrimaryColor)
		case 2: // 成功
			icon = "✓"
			style = lipgloss.NewStyle().Foreground(theme.SuccessColor)
		case 3: // 失败
			icon = "✗"
			style = lipgloss.NewStyle().Foreground(theme.ErrorColor)
		}

		s.WriteString(fmt.Sprintf("  %s %s\n",
			style.Render(icon),
			style.Render(fmt.Sprintf("Step %d: %s", i+1, cmd.Description))))
	}

	// 错误信息
	if m.done && m.err != nil {
		s.WriteString("\n")
		errStyle := lipgloss.NewStyle().Foreground(theme.ErrorColor).Bold(true)
		s.WriteString(errStyle.Render("❌ Error: " + m.err.Error()))
		s.WriteString("\n")
	}

	// 完成信息
	if m.done {
		s.WriteString("\n")
		if m.err != nil {
			s.WriteString(lipgloss.NewStyle().
				Foreground(theme.ErrorColor).
				Bold(true).
				Render("Process failed"))
		}
	}

	return s.String()
}

// 执行当前命令
func (m *CommandStepModel) runCommand() tea.Cmd {
	if m.currentStep >= len(m.Commands) {
		return nil
	}

	// 标记当前步骤为运行中
	m.stepStatus[m.currentStep] = 1

	return func() tea.Msg {
		cmd := m.Commands[m.currentStep]
		execCmd := exec.Command(cmd.Command, cmd.Args...)
		output, err := execCmd.CombinedOutput()

		return stepCompleteMsg{
			index:  m.currentStep,
			err:    err,
			output: string(output),
		}
	}
}

// RunMultipleCommandsWithSimpleProgress 执行多个命令并使用 Bubble Tea 显示进度
func RunMultipleCommandsWithSimpleProgress(commands []CommandInfo) error {
	model := NewCommandStepModel(commands)
	p := tea.NewProgram(model)

	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("failed to run progress UI: %w", err)
	}

	// 检查最终状态
	if m, ok := finalModel.(*CommandStepModel); ok {
		if m.err != nil {
			return m.err
		}
		// 所有命令成功后，直接输出成功提示到屏幕
		fmt.Println("🎉 All commands completed successfully!")
	}

	return nil
}
