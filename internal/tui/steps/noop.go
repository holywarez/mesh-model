package steps

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

const BRAILLE_OFFSET = '\u2800'

var BRAILLE = [4][2]rune{
	{'\u0001', '\u0008'},
	{'\u0002', '\u0010'},
	{'\u0004', '\u0020'},
	{'\u0040', '\u0080'},
}

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
	x := BRAILLE[1][0] | BRAILLE[1][1] + BRAILLE_OFFSET
	y := BRAILLE[2][0] | BRAILLE[2][1] + BRAILLE_OFFSET
	return fmt.Sprintf("Noop Step: %s", string([]rune{x, y}))
}

func (t Noop) SelectedValueText() string {
	return ""
}
