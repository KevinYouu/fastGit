package spinner

import (
	"fmt"
	"strings"
	"time"

	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// AdvancedSpinnerType 高级加载动画类型
type AdvancedSpinnerType int

const (
	SpinnerDefault AdvancedSpinnerType = iota
	SpinnerPulse
	SpinnerDots
	SpinnerArrow
	SpinnerProgress
)

// AdvancedSpinnerModel 高级加载动画模型
type AdvancedSpinnerModel struct {
	spinner      spinner.Model
	spinnerType  AdvancedSpinnerType
	message      string
	submessage   string
	progress     float64
	steps        []string
	currentStep  int
	done         bool
	success      bool
	err          error
	resultMsg    string
	showProgress bool
	showSteps    bool
	elapsedTime  time.Duration
	startTime    time.Time
}

// NewAdvancedSpinner 创建新的高级加载动画
func NewAdvancedSpinner(spinnerType AdvancedSpinnerType, message string) AdvancedSpinnerModel {
	s := spinner.New()

	// 根据类型设置不同的动画帧
	switch spinnerType {
	case SpinnerPulse:
		s.Spinner = spinner.Spinner{
			Frames: theme.GetPulseSpinnerFrames(),
			FPS:    time.Millisecond * 150,
		}
	case SpinnerDots:
		s.Spinner = spinner.Spinner{
			Frames: theme.GetDotsSpinnerFrames(),
			FPS:    time.Millisecond * 100,
		}
	case SpinnerArrow:
		s.Spinner = spinner.Spinner{
			Frames: theme.GetArrowSpinnerFrames(),
			FPS:    time.Millisecond * 200,
		}
	default:
		s.Spinner = spinner.Spinner{
			Frames: theme.GetSpinnerFrames(),
			FPS:    time.Millisecond * 100,
		}
	}

	s.Style = theme.GetSpinnerStyle()

	return AdvancedSpinnerModel{
		spinner:     s,
		spinnerType: spinnerType,
		message:     message,
		startTime:   time.Now(),
	}
}

// SetMessage 设置主消息
func (m *AdvancedSpinnerModel) SetMessage(message string) {
	m.message = message
}

// SetSubmessage 设置子消息
func (m *AdvancedSpinnerModel) SetSubmessage(submessage string) {
	m.submessage = submessage
}

// SetProgress 设置进度
func (m *AdvancedSpinnerModel) SetProgress(progress float64) {
	m.progress = progress
	m.showProgress = true
}

// SetSteps 设置步骤列表
func (m *AdvancedSpinnerModel) SetSteps(steps []string) {
	m.steps = steps
	m.showSteps = true
}

// NextStep 前进到下一步
func (m *AdvancedSpinnerModel) NextStep() {
	if m.currentStep < len(m.steps)-1 {
		m.currentStep++
	}
}

// SetDone 设置完成状态
func (m *AdvancedSpinnerModel) SetDone(success bool, resultMsg string, err error) {
	m.done = true
	m.success = success
	m.resultMsg = resultMsg
	m.err = err
	m.elapsedTime = time.Since(m.startTime)
}

// Init 初始化
func (m AdvancedSpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update 更新
func (m AdvancedSpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.done {
		return m, nil
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// View 渲染
func (m AdvancedSpinnerModel) View() string {
	if m.done {
		return m.renderComplete()
	}

	var content strings.Builder

	// 渲染标题区域
	content.WriteString(m.renderHeader())
	content.WriteString("\n")

	// 渲染主消息区域
	content.WriteString(m.renderMainMessage())
	content.WriteString("\n")

	// 渲染进度条（如果启用）
	if m.showProgress {
		content.WriteString(m.renderProgress())
		content.WriteString("\n")
	}

	// 渲染步骤列表（如果启用）
	if m.showSteps {
		content.WriteString(m.renderSteps())
		content.WriteString("\n")
	}

	// 渲染子消息
	if m.submessage != "" {
		content.WriteString(m.renderSubmessage())
		content.WriteString("\n")
	}

	return content.String()
}

// renderHeader 渲染标题
func (m AdvancedSpinnerModel) renderHeader() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Background(theme.BackgroundHighlight).
		Bold(true).
		Padding(1, 2).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(theme.PrimaryColor).
		Width(60).
		Align(lipgloss.Center)

	return headerStyle.Render("🚀 FastGit 操作进行中...")
}

// renderMainMessage 渲染主消息
func (m AdvancedSpinnerModel) renderMainMessage() string {
	iconStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Bold(true)

	messageStyle := lipgloss.NewStyle().
		Foreground(theme.TextColor).
		Bold(true).
		MarginLeft(2)

	return fmt.Sprintf("%s %s",
		iconStyle.Render(m.spinner.View()),
		messageStyle.Render(m.message))
}

// renderProgress 渲染进度条
func (m AdvancedSpinnerModel) renderProgress() string {
	width := 40
	filled := int(m.progress * float64(width))

	var progressBar strings.Builder
	for i := 0; i < width; i++ {
		if i < filled {
			progressBar.WriteString("█")
		} else {
			progressBar.WriteString("░")
		}
	}

	progressStyle := lipgloss.NewStyle().
		Foreground(theme.PrimaryColor).
		Background(theme.BackgroundLighter).
		Padding(0, 1)

	percentStyle := lipgloss.NewStyle().
		Foreground(theme.AccentColor).
		Bold(true).
		MarginLeft(2)

	return fmt.Sprintf("  %s %s",
		progressStyle.Render(progressBar.String()),
		percentStyle.Render(fmt.Sprintf("%.1f%%", m.progress*100)))
}

// renderSteps 渲染步骤列表
func (m AdvancedSpinnerModel) renderSteps() string {
	var steps strings.Builder
	steps.WriteString("  步骤进度:\n")

	for i, step := range m.steps {
		var icon, style string
		if i < m.currentStep {
			icon = theme.GetStatusIcon("success")
			style = theme.SuccessStyle.Render(step)
		} else if i == m.currentStep {
			icon = theme.GetStatusIcon("loading")
			style = theme.InfoStyle.Render(step)
		} else {
			icon = theme.GetStatusIcon("pending")
			style = theme.UnselectedStyle.Render(step)
		}

		steps.WriteString(fmt.Sprintf("    %s %s\n", icon, style))
	}

	return steps.String()
}

// renderSubmessage 渲染子消息
func (m AdvancedSpinnerModel) renderSubmessage() string {
	submessageStyle := lipgloss.NewStyle().
		Foreground(theme.TextSecondary).
		Italic(true).
		MarginLeft(4)

	return submessageStyle.Render(fmt.Sprintf("💡 %s", m.submessage))
}

// renderComplete 渲染完成状态
func (m AdvancedSpinnerModel) renderComplete() string {
	var content strings.Builder

	// 渲染完成标题
	var headerStyle lipgloss.Style
	var icon string
	var title string

	if m.success {
		headerStyle = lipgloss.NewStyle().
			Foreground(theme.TextColor).
			Background(theme.SuccessColor).
			Bold(true).
			Padding(1, 2).
			Width(60).
			Align(lipgloss.Center)
		icon = "✨"
		title = "操作成功完成！"
	} else {
		headerStyle = lipgloss.NewStyle().
			Foreground(theme.TextColor).
			Background(theme.ErrorColor).
			Bold(true).
			Padding(1, 2).
			Width(60).
			Align(lipgloss.Center)
		icon = "❌"
		title = "操作失败"
	}

	content.WriteString(headerStyle.Render(fmt.Sprintf("%s %s", icon, title)))
	content.WriteString("\n\n")

	// 渲染结果消息
	resultStyle := lipgloss.NewStyle().
		Foreground(theme.TextColor).
		Bold(true).
		Padding(1, 2).
		Background(theme.BackgroundHighlight)

	content.WriteString(resultStyle.Render(m.resultMsg))

	// 如果有错误，显示错误信息
	if m.err != nil {
		content.WriteString("\n")
		errorStyle := lipgloss.NewStyle().
			Foreground(theme.ErrorColor).
			Italic(true).
			Padding(0, 2)
		content.WriteString(errorStyle.Render(fmt.Sprintf("错误详情: %v", m.err)))
	}

	// 显示耗时
	content.WriteString("\n")
	timeStyle := lipgloss.NewStyle().
		Foreground(theme.TextMuted).
		Italic(true).
		Padding(0, 2)
	content.WriteString(timeStyle.Render(fmt.Sprintf("⏱️ 耗时: %v", m.elapsedTime.Round(time.Millisecond))))

	return content.String()
}

// MultiStepSpinner 多步骤加载动画
type MultiStepSpinner struct {
	steps         []string
	currentStep   int
	spinner       AdvancedSpinnerModel
	stepDurations []time.Duration
	stepMessages  []string
}

// NewMultiStepSpinner 创建多步骤加载动画
func NewMultiStepSpinner(steps []string, stepMessages []string) *MultiStepSpinner {
	spinner := NewAdvancedSpinner(SpinnerDefault, "初始化...")
	spinner.SetSteps(steps)

	return &MultiStepSpinner{
		steps:        steps,
		spinner:      spinner,
		stepMessages: stepMessages,
	}
}

// ExecuteStep 执行步骤
func (m *MultiStepSpinner) ExecuteStep(stepIndex int, message string) {
	if stepIndex < len(m.steps) {
		m.currentStep = stepIndex
		m.spinner.SetMessage(message)
		m.spinner.currentStep = stepIndex

		// 更新进度
		progress := float64(stepIndex) / float64(len(m.steps))
		m.spinner.SetProgress(progress)
	}
}

// Complete 完成所有步骤
func (m *MultiStepSpinner) Complete(success bool, message string, err error) {
	m.spinner.SetProgress(1.0)
	m.spinner.SetDone(success, message, err)
}

// GetSpinner 获取内部的 spinner 用于 tea 程序
func (m *MultiStepSpinner) GetSpinner() AdvancedSpinnerModel {
	return m.spinner
}
