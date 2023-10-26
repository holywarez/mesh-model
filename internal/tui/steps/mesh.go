package steps

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	label string
	desc  string
}

func (i item) Title() string       { return i.label }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.label }

type meshModel struct {
	list          list.Model
	selectedValue string
}

func NewMeshStep() WizardStep {
	items := []list.Item{
		item{
			label: "Mesh 1",
			desc:  "This is the first mesh",
		},
		item{
			label: "Mesh 2",
			desc:  "This is the second mesh",
		},
		item{
			label: "Mesh 3",
			desc:  "This is the third mesh",
		},
		item{
			label: "Add Mesh",
			desc:  "Add a new mesh",
		},
	}

	return meshModel{
		list: list.New(
			items,
			list.NewDefaultDelegate(),
			30, 30,
		),
	}
}

func (m meshModel) Init() tea.Cmd {
	return nil
}

func (m meshModel) ResetList() meshModel {
	m.list.SetShowPagination(false)
	m.list.SetShowHelp(false)
	m.list.SetShowFilter(false)
	m.list.SetFilteringEnabled(false)
	m.list.SetShowStatusBar(false)
	m.list.SetShowTitle(false)
	return m
}

func (m meshModel) Update(msg tea.Msg) (WizardStep, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc", "left":
			return m, ReturnCmd
		case "enter", "right":
			m.selectedValue = m.list.SelectedItem().(item).Title()
			return m, ReturnCmd
		}

	case tea.WindowSizeMsg:
		_, h := docStyle.GetFrameSize()
		m.list.SetSize(30, msg.Height-12-h)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m meshModel) View() string {
	model := m.ResetList()
	return docStyle.Render(model.list.View())
}

func (m meshModel) Filled() bool {
	return m.selectedValue != ""
}

func (m meshModel) SelectedValueText() string {
	return m.selectedValue
}
