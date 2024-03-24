package input

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Input function starts a new input program with a provided prompt, placeholder,and exit prompt.
func Input(prompt, placeholder, exitPrompt string) string {
	p := tea.NewProgram(initialModel(prompt, placeholder, exitPrompt))
	if data, err := p.Run(); err != nil {
		log.Fatal(err)
		return ""
	} else {
		model := data.(model)
		userInput := model.textInput.Value()
		return userInput
	}
}

type (
	errMsg error
)

type model struct {
	textInput  textinput.Model
	err        error
	prompt     string // New field to store the prompt
	quitPrompt string
	quitting   bool
}

// initialModel function initializes a new model with the provided prompt, placeholder,and exit prompt.
func initialModel(prompt, placeholder, quitPrompt string) model {
	ti := textinput.New()
	ti.Placeholder = placeholder // Set input placeholder
	ti.Focus()                   // Set input focus
	ti.CharLimit = 156           // Set input character limit
	ti.Width = 60                // Set input width

	// Return the initialized model, including the input model and prompt
	return model{
		textInput:  ti,
		err:        nil,
		prompt:     prompt,
		quitPrompt: quitPrompt,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		m.quitting = true
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View method renders the user interface, including the prompt and input field
func (m model) View() string {
	if m.quitting {
		return "" // clear the screen
	}

	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.prompt,           // Display the prompt
		m.textInput.View(), // Display the input field
		m.quitPrompt,       // Display the exit prompt
	) + "\n"
}
