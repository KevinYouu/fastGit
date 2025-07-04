package spinner

import (
	"fmt"
	"time"

	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// SpinnerModel 加载动画模型
type SpinnerModel struct {
	spinner   spinner.Model
	message   string
	done      bool
	success   bool
	err       error
	resultMsg string
}

// NewSpinner 创建新的加载动画
func NewSpinner(message string) SpinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Spinner{
		Frames: theme.GetSpinnerFrames(),
		FPS:    time.Millisecond * 100,
	}
	s.Style = theme.GetSpinnerStyle()

	return SpinnerModel{
		spinner: s,
		message: message,
	}
}

// SetMessage 设置消息
func (m *SpinnerModel) SetMessage(message string) {
	m.message = message
}

// SetDone 设置完成状态
func (m *SpinnerModel) SetDone(success bool, resultMsg string, err error) {
	m.done = true
	m.success = success
	m.resultMsg = resultMsg
	m.err = err
}

// Init 初始化
func (m SpinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update 更新
func (m SpinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

// View 渲染
func (m SpinnerModel) View() string {
	if m.done {
		if m.success {
			return fmt.Sprintf("%s %s\n",
				theme.SuccessStyle.Render("✓"),
				theme.SuccessStyle.Render(m.resultMsg))
		} else {
			errorMsg := m.resultMsg
			if m.err != nil {
				errorMsg = fmt.Sprintf("%s: %v", m.resultMsg, m.err)
			}
			return fmt.Sprintf("%s %s\n",
				theme.ErrorStyle.Render("✗"),
				theme.ErrorStyle.Render(errorMsg))
		}
	}

	return fmt.Sprintf("%s %s",
		m.spinner.View(),
		theme.InfoStyle.Render(m.message))
}

// SimpleSpinner 简单的加载动画函数
func SimpleSpinner(message string, duration time.Duration) {
	frames := theme.GetSpinnerFrames()
	style := theme.GetSpinnerStyle()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)
	go func() {
		time.Sleep(duration)
		done <- true
	}()

	frameIndex := 0
	for {
		select {
		case <-done:
			// 清除当前行并显示完成消息
			fmt.Printf("\r%s %s\n",
				theme.SuccessStyle.Render("✓"),
				theme.SuccessStyle.Render(message))
			return
		case <-ticker.C:
			// 显示当前帧
			fmt.Printf("\r%s %s",
				style.Render(frames[frameIndex]),
				theme.InfoStyle.Render(message))
			frameIndex = (frameIndex + 1) % len(frames)
		}
	}
}
