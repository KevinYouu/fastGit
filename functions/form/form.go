package form

import (
	"fmt"
	"log"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// "Enter the following information:",
// inputStyle.Width(11).Render("Remote name"),
// m.inputs[name].View(),
// inputStyle.Width(11).Render("Remote url"),

type FormProps struct {
	Message      string
	Field        string
	Field2       string
	FieldLength  int
	Field2Length int
}

func FormInput(props FormProps) (string, string, error) {
	p := tea.NewProgram(initialModel(props))

	if data, err := p.Run(); err != nil {
		log.Fatal(err)
		return "", "", err
	} else {
		inputs := data.(model).inputs
		return inputs[0].Value(), inputs[1].Value(), nil
	}
}

type (
	errMsg error
)

const (
	name = iota
	url
)

const (
	hotPink = lipgloss.Color("#0091FF")
)

var (
	inputStyle = lipgloss.NewStyle().Foreground(hotPink)
	// continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

type model struct {
	inputs  []textinput.Model
	focused int
	props   FormProps
	err     error
}

// Validator functions to ensure valid input
func nameValidator(s string) error {
	// Name should not be empty
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func urlValidator(s string) error {
	// URL should be a valid URL
	if len(strings.TrimSpace(s)) == 0 {
		return fmt.Errorf("URL cannot be empty")
	}
	// For simplicity, we're not checking for a valid URL format here
	return nil
}

func initialModel(props FormProps) model {
	var inputs []textinput.Model = make([]textinput.Model, 2)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "origin"
	inputs[name].Focus()
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	inputs[name].Validate = nameValidator

	inputs[url] = textinput.New()
	inputs[url].Placeholder = "https://github.com/KevinYouu/fastGit"
	inputs[url].Width = 50
	inputs[url].Prompt = ""
	inputs[url].Validate = urlValidator

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
		props:   props,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(`%s
%s: %s
%s: %s
`,
		colors.RenderColor("white", m.props.Message),
		inputStyle.Width(m.props.FieldLength).Render(m.props.Field),
		m.inputs[name].View(),
		inputStyle.Width(m.props.Field2Length).Render(m.props.Field2),
		m.inputs[url].View(),
	)
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}
