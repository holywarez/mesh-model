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
	steps   []steps.StepDefinition
	cursor  int
	active  int
	width   int
	height  int
}

var (
	menuBaseStyle     = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Padding(1, 3, 1, 3)
	menuActiveStyle   = menuBaseStyle.Copy().BorderForeground(lipgloss.Color("105"))
	menuInactiveStyle = menuBaseStyle.Copy().BorderForeground(lipgloss.Color("235"))
	hintStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Padding(1)
)

func New() TUI {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return TUI{
		spinner: s,
		active:  -1,
		width:   1,
		height:  1,
		steps: []steps.StepDefinition{
			steps.NewStepDefinition("0. Render Colors", steps.NewColorsStep()),
			steps.NewStepDefinition("1. Select Mesh", steps.NewMeshStep()),
			steps.NewStepDefinition("2. Select Algo", steps.NewNoopStep()),
			steps.NewStepDefinition("3. Run Simulation", steps.NewNoopStep()),
		},
	}
}

func (t TUI) Init() tea.Cmd {
	var cmds []tea.Cmd

	for _, step := range t.steps {
		cmds = append(cmds, step.Step.Init())
	}

	cmds = append(cmds, t.spinner.Tick)

	return tea.Batch(cmds...)
}

func (t TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil
	switch msg := msg.(type) {

	case steps.ReturnMsg:
		t.active = -1

	case tea.WindowSizeMsg:
		t.width = msg.Width
		t.height = msg.Height - 10
		if t.active >= 0 {
			t.steps[t.active].Step, cmd = t.steps[t.active].Step.Update(msg)
			return t, cmd
		}

	case tea.KeyMsg:
		if t.active >= 0 {
			step := t.steps[t.active].Step
			t.steps[t.active].Step, cmd = step.Update(msg)
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
		case "enter", " ", "right":
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
	var views []string

	w := fmt.Sprintf("%s Setup Mesh Network Simulation\n\n", t.spinner.View())
	s := ""

	for i, step := range t.steps {
		s += step.View(i == t.cursor) + "\n"
	}

	widthStyle := lipgloss.NewStyle().Width(30)
	heightStyle := lipgloss.NewStyle().Height(t.height)

	if t.active >= 0 {
		views = append(views, menuInactiveStyle.Copy().Inherit(widthStyle).Inherit(heightStyle).Render(s))
		views = append(views, menuActiveStyle.Copy().Inherit(heightStyle).Render(t.steps[t.active].Step.View()))
	} else {
		views = append(views, menuActiveStyle.Copy().Inherit(widthStyle).Inherit(heightStyle).Render(s))
	}

	w += lipgloss.JoinHorizontal(lipgloss.Top, views...)
	// The footer
	w += hintStyle.Render("\nPress q to quit.\n")

	// Send the UI for rendering
	return w
}
