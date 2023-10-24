package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"mesh-model/internal/tui/steps"
)

type TUI struct {
	spinner spinner.Model
	steps   []StepDefinition
	cursor  int
	active  int
}

type StepDefinition struct {
	label string
	step  steps.WizardStep
}

func New() TUI {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return TUI{
		spinner: s,
		steps: []StepDefinition{
			{label: "1. Select Mesh", step: steps.NewNoopStep()},
			{label: "2. Select Algo", step: steps.NewNoopStep()},
			{label: "3. Run Simulation", step: steps.NewNoopStep()},
		},
	}
}

func (t TUI) Init() tea.Cmd {
	return t.spinner.Tick
}

func (t TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the prograt.
		case "ctrl+c", "q":
			return t, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if t.cursor > 0 {
				t.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if t.cursor < len(t.steps)-1 {
				t.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := t.steps[t.cursor]
			if ok {
				delete(t.steps, t.cursor)
			} else {
				t.steps[t.cursor] = struct{}{}
			}
		}
	default:
		var cmd tea.Cmd
		t.spinner, cmd = t.spinner.Update(msg)
		return t, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return t, nil
}

func (t TUI) View() string {
	s := fmt.Sprintf("%s Setup Mesh Network Simulation\n\n", t.spinner.View())

	for i, step := range t.steps {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if t.cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := t.selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
