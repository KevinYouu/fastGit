package confirm

import (
	"log"

	"github.com/KevinYouu/fastGit/functions/colors"

	tea "github.com/charmbracelet/bubbletea"
)

func Confirm(prompt string) bool {
	p := tea.NewProgram(initialModel(prompt))

	if data, err := p.Run(); err != nil {
		log.Fatal(err)
		return false
	} else {
		value := data.(model).cursor
		return value
	}
}

type model struct {
	cursor   bool
	quitting bool
	prompt   string
}

func initialModel(prompt string) model {
	if prompt == "" {
		prompt = "Are you sure?"
	}
	return model{cursor: true, prompt: prompt}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left", "h":
			m.cursor = true
		case "right", "l":
			m.cursor = false
		case "enter", " ":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	yesStyle := "[ ]"
	noStyle := "[ ]"
	if m.cursor {
		yesStyle = colors.RenderColor("blue", "[x] yes")
		noStyle = colors.RenderColor("white", "[ ] no")
	} else {
		yesStyle = colors.RenderColor("white", "[ ] yes")
		noStyle = colors.RenderColor("red", "[x] no")
	}

	if m.quitting {
		return "" // clear the screen
	}
	return m.prompt + "\n\n" + yesStyle + "    " + noStyle +
		"\n\n(use left/right arrow keys to choose, enter to confirm)"
}
