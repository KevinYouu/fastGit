package form

import (
	"fmt"
	"io"
	"os"

	"github.com/KevinYouu/fastGit/internal/config"
	"github.com/KevinYouu/fastGit/internal/theme"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type item struct {
	title string
	value string
}

func (i item) FilterValue() string { return i.title }

type itemDelegate struct {
	compact bool
}

func (d itemDelegate) Height() int {
	if d.compact {
		return 1 // 紧凑模式下每项占用1行
	}
	return 2 // 普通模式下每项占用2行
}

func (d itemDelegate) Spacing() int {
	if d.compact {
		return 0 // 紧凑模式下无间距
	}
	return 1 // 普通模式下有间距
}

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	// 根据模式决定显示格式
	var str string
	if d.compact {
		// 紧凑模式：单行显示，截断长文本
		str = truncateText(i.title, 60)
	} else {
		// 普通模式：保持原格式
		str = i.title
	}

	fn := lipgloss.NewStyle().
		Foreground(theme.TextColor).
		PaddingLeft(1).
		Render

	if index == m.Index() {
		fn = lipgloss.NewStyle().
			Foreground(theme.TextColor).
			Background(theme.BackgroundActive).
			Bold(true).
			PaddingLeft(1).
			PaddingRight(1).
			Render
	}

	fmt.Fprint(w, fn(str))
}

// 截断文本，保留重要信息
func truncateText(text string, maxLength int) string {
	if len(text) <= maxLength {
		return text
	}

	// 对于多行文本，只保留第一行
	lines := []string{}
	current := ""
	for _, char := range text {
		if char == '\n' {
			if current != "" {
				lines = append(lines, current)
				current = ""
			}
		} else {
			current += string(char)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}

	if len(lines) > 0 {
		firstLine := lines[0]
		if len(firstLine) > maxLength {
			return firstLine[:maxLength-3] + "..."
		}
		return firstLine
	}

	return text[:maxLength-3] + "..."
}

type noLoopModel struct {
	list     list.Model
	choice   string
	quitting bool
	compact  bool
}

func (m noLoopModel) Init() tea.Cmd {
	return nil
}

func (m noLoopModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = i.value
			}
			return m, tea.Quit

		case "up", "k":
			// 禁用循环：如果已经在第一项，不移动
			if m.list.Index() > 0 {
				m.list.CursorUp()
			}
			return m, nil

		case "down", "j":
			// 禁用循环：如果已经在最后一项，不移动
			if m.list.Index() < len(m.list.Items())-1 {
				m.list.CursorDown()
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m noLoopModel) View() string {
	if m.choice != "" {
		return ""
	}
	if m.quitting {
		return ""
	}

	if m.compact {
		// 紧凑模式：去掉额外的换行符
		return m.list.View()
	} else {
		// 普通模式：保持原有格式
		return "\n" + m.list.View()
	}
}

// NoLoopSelectForm 创建一个不循环的选择表单，支持自适应终端高度
func NoLoopSelectForm(options []config.Option) (label, value string, err error) {
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = item{title: opt.Label, value: opt.Value}
	}

	// 检测终端尺寸
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// 如果无法获取终端尺寸，使用默认值
		width = 80
		height = 24
	}

	// 根据终端高度决定是否使用紧凑模式
	compact := height < 15

	// 计算可用的列表高度
	var listHeight int
	if compact {
		// 紧凑模式：减少标题、帮助信息等的空间占用
		listHeight = height - 4 // 预留标题和帮助信息的空间
		if listHeight < 3 {
			listHeight = 3
		}
	} else {
		// 普通模式：保留更多空间给标题和帮助信息
		listHeight = height - 8
		if listHeight < 5 {
			listHeight = 5
		}
	}

	// 创建列表，使用自适应的委托
	delegate := itemDelegate{compact: compact}
	l := list.New(items, delegate, width, listHeight)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false) // 不显示帮助信息

	// 根据模式调整样式
	if compact {
		// 紧凑模式：简化标题样式
		l.Styles.Title = lipgloss.NewStyle().
			Foreground(theme.PrimaryColor).
			Bold(true)
		l.Styles.HelpStyle = lipgloss.NewStyle().
			Foreground(theme.TextSecondary).
			Height(1)
	} else {
		// 普通模式：保持原有样式
		l.Styles.Title = lipgloss.NewStyle().
			Foreground(theme.PrimaryColor).
			Bold(true).
			Padding(0, 1)
		l.Styles.HelpStyle = lipgloss.NewStyle().
			Foreground(theme.TextSecondary)
	}

	// 自定义按键映射，移除默认的循环行为
	l.KeyMap.CursorUp = key.NewBinding(
		key.WithKeys("up", "k"),
	)
	l.KeyMap.CursorDown = key.NewBinding(
		key.WithKeys("down", "j"),
	)

	m := noLoopModel{list: l, compact: compact}

	// 根据模式选择是否使用全屏
	var p *tea.Program
	if compact {
		// 紧凑模式：不使用全屏，减少空间占用
		p = tea.NewProgram(m)
	} else {
		// 普通模式：使用全屏
		p = tea.NewProgram(m, tea.WithAltScreen())
	}

	finalModel, err := p.Run()
	if err != nil {
		return "", "", err
	}

	if m, ok := finalModel.(noLoopModel); ok {
		if m.quitting && m.choice == "" {
			return "", "", fmt.Errorf("user aborted")
		}
		// 找到选中的选项
		for _, opt := range options {
			if opt.Value == m.choice {
				return opt.Label, opt.Value, nil
			}
		}
	}

	return "", "", fmt.Errorf("no selection made")
}

// NoLoopSelectFormWithStringSlice 创建一个基于字符串切片的不循环选择表单，支持自适应终端高度
func NoLoopSelectFormWithStringSlice(title string, options []string) (label, value string, err error) {
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = item{title: opt, value: opt}
	}

	// 检测终端尺寸
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		// 如果无法获取终端尺寸，使用默认值
		width = 80
		height = 24
	}

	// 根据终端高度决定是否使用紧凑模式
	compact := height < 15

	// 计算可用的列表高度
	var listHeight int
	if compact {
		// 紧凑模式：减少标题、帮助信息等的空间占用
		listHeight = height - 4 // 预留标题和帮助信息的空间
		if listHeight < 3 {
			listHeight = 3
		}
	} else {
		// 普通模式：保留更多空间给标题和帮助信息
		listHeight = height - 8
		if listHeight < 5 {
			listHeight = 5
		}
	}

	// 创建列表，使用自适应的委托
	delegate := itemDelegate{compact: compact}
	l := list.New(items, delegate, width, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false) // 不显示帮助信息

	// 根据模式调整样式
	if compact {
		// 紧凑模式：简化标题样式
		l.Styles.Title = lipgloss.NewStyle().
			Foreground(theme.PrimaryColor).
			Bold(true)
		l.Styles.HelpStyle = lipgloss.NewStyle().
			Foreground(theme.TextSecondary).
			Height(1)
	} else {
		// 普通模式：保持原有样式
		l.Styles.Title = lipgloss.NewStyle().
			Foreground(theme.PrimaryColor).
			Bold(true).
			Padding(0, 1)
		l.Styles.HelpStyle = lipgloss.NewStyle().
			Foreground(theme.TextSecondary)
	}

	// 自定义按键映射，移除默认的循环行为
	l.KeyMap.CursorUp = key.NewBinding(
		key.WithKeys("up", "k"),
	)
	l.KeyMap.CursorDown = key.NewBinding(
		key.WithKeys("down", "j"),
	)

	m := noLoopModel{list: l, compact: compact}

	// 根据模式选择是否使用全屏
	var p *tea.Program
	if compact {
		// 紧凑模式：不使用全屏，减少空间占用
		p = tea.NewProgram(m)
	} else {
		// 普通模式：使用全屏
		p = tea.NewProgram(m, tea.WithAltScreen())
	}

	finalModel, err := p.Run()
	if err != nil {
		return "", "", err
	}

	if m, ok := finalModel.(noLoopModel); ok {
		if m.quitting && m.choice == "" {
			return "", "", fmt.Errorf("user aborted")
		}
		return m.choice, m.choice, nil
	}

	return "", "", fmt.Errorf("no selection made")
}
