package steps

import tea "github.com/charmbracelet/bubbletea"

type meshModel struct {
}

func NewMeshStep() WizardStep {
	return meshModel{}
}

func (m meshModel) Init() tea.Cmd {
	return nil
}

func (m meshModel) Update(msg tea.Msg) (WizardStep, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, ReturnCmd
		}
	}

	return m, nil
}

func (m meshModel) View() string {
	return "Select Mesh Model"
}

func (m meshModel) Filled() bool {
	return false
}
