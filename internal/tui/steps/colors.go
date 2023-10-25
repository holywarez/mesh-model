package steps

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct{}

func NewColorsStep() WizardStep {
	return model{}
}

func (m model) Filled() bool {
	return true
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (WizardStep, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, ReturnCmd
		}
	}

	return m, nil
}

func (m model) View() string {
	var s = ""

	for i := 0; i < 256; i++ {
		s += lipgloss.
			NewStyle().Foreground(lipgloss.Color(fmt.Sprintf("%d", i))).
			Render(fmt.Sprintf("%d color \t", i))
		if i%10 == 0 {
			s += "\n"
		}
	}

	return lipgloss.NewStyle().Padding(1, 3).Render(s)
}
