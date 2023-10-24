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

var (
	selected_style    = lipgloss.NewStyle().Foreground(lipgloss.Color("255"))
	filled_step_style = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	empty_step_style  = lipgloss.NewStyle().Foreground(lipgloss.Color("235"))
)

func New() TUI {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return TUI{
		spinner: s,
		active:  -1,
		steps: []StepDefinition{
			{label: "0. Render Colors", step: steps.NewColorsStep()},
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
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {

	case steps.ReturnMsg:
		t.active = -1
	case tea.KeyMsg:
		if t.active >= 0 {
			step := t.steps[t.active].step
			t.steps[t.active].step, cmd = step.Update(msg)
			return t, cmd
		}

		switch msg.String() {

		case "ctrl+c", "q":
			return t, tea.Quit

		case "up", "k":
			if t.cursor > 0 {
				t.cursor--
			}

		case "down", "j":
			if t.cursor < len(t.steps)-1 {
				t.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			t.active = t.cursor
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
	if t.active >= 0 {
		return t.steps[t.active].step.View()
	}

	s := fmt.Sprintf("%s Setup Mesh Network Simulation\n\n", t.spinner.View())

	for i, step := range t.steps {
		cursor := " " // no cursor
		if t.cursor == i {
			cursor = "|" // cursor!
		}

		var style *lipgloss.Style = nil
		if t.cursor == i {
			style = &selected_style
		} else {
			if step.step.Filled() {
				style = &filled_step_style
			} else {
				style = &empty_step_style
			}
		}
		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, style.Render(step.label))
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
