package multipleChoice

import (
	"fmt"
	"os"
	"strings"

	"github.com/KevinYouu/fastGit/functions/colors"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor   int
	choices  []string // Choices passed as a parameter
	selected map[int]bool
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			os.Exit(1)
			return m, tea.Quit
		case "enter":
			m.quitting = true
			return m, tea.Quit

		case " ":
			// Toggle selection on enter or space press
			m.selected[m.cursor] = !m.selected[m.cursor]

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("What kind of Bubble Tea would you like to order?\n")

	for i := 0; i < len(m.choices); i++ {
		if m.selected[i] {
			s.WriteString(colors.RenderColor("green", "✓"))
		} else {
			s.WriteString("○")
		}
		s.WriteString(" ")

		if m.cursor == i {
			s.WriteString(colors.RenderColor("blue", m.choices[i]))
		} else {
			if m.selected[i] {
				s.WriteString(colors.RenderColor("green", m.choices[i]))
			} else {
				s.WriteString(colors.RenderColor("white", m.choices[i]))
			}
		}
		s.WriteString("\n")
	}
	s.WriteString("\x1b[0m") // reset color
	s.WriteString("(press q to quit)\n")
	if m.quitting {
		return "" // clear the screen
	}
	return s.String()
}

func MultipleChoice(choices []string) []string {
	initialSelected := make(map[int]bool)
	for range choices {
		initialSelected[len(initialSelected)] = false
	}

	p := tea.NewProgram(model{
		choices:  choices,
		selected: initialSelected,
	})

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok {
		var selectedChoices []string
		for i, isSelected := range m.selected {
			if isSelected {
				selectedChoices = append(selectedChoices, m.choices[i])
			}
		}
		return selectedChoices
	}

	return nil
}
