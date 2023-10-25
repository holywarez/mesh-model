package steps

import tea "github.com/charmbracelet/bubbletea"

type WizardStep interface {
	Filled() bool
	Init() tea.Cmd
	Update(msg tea.Msg) (WizardStep, tea.Cmd)
	View() string
}

func ReturnCmd() tea.Msg {
	return ReturnMsg{}
}

type ReturnMsg struct{}
