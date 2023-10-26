package steps

import tea "github.com/charmbracelet/bubbletea"

type Noop struct {
}

func NewNoopStep() Noop {
	return Noop{}
}

func (t Noop) Filled() bool {
	return false
}

func (t Noop) Init() tea.Cmd {
	return nil
}

func (t Noop) Update(msg tea.Msg) (WizardStep, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "left":
			return t, ReturnCmd
		}
	}

	return t, nil
}

func (t Noop) View() string {
	return "Noop Step"
}

func (t Noop) SelectedValueText() string {
	return ""
}
