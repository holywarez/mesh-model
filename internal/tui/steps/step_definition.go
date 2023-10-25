package steps

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

type StepDefinition struct {
	Label string
	Step  WizardStep
}

func NewStepDefinition(label string, step WizardStep) StepDefinition {
	return StepDefinition{
		Label: label,
		Step:  step,
	}
}

var (
	menuItemStyle            = lipgloss.NewStyle().Padding(1, 0)
	menuCaptionStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))
	menuSelectedCaptionStyle = menuCaptionStyle.Copy().Bold(true).Foreground(lipgloss.Color("255"))
	menuValueStyle           = lipgloss.NewStyle().MaxWidth(24).Foreground(lipgloss.Color("245"))
	menuCursorStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("135"))
)

func (sd StepDefinition) View(selected bool) string {
	var (
		caret        = " "
		captionStyle = menuSelectedCaptionStyle.Copy().Faint(!selected)
		valueStyle   = menuValueStyle.Copy().Faint(!selected)
	)

	if selected {
		caret = menuCursorStyle.Render("│")
	}

	rows := []string{
		fmt.Sprintf("%s %s", caret, captionStyle.Render(sd.Label)),
		fmt.Sprintf("%s %s", caret, valueStyle.Render(sd.ViewValue())),
	}

	return menuItemStyle.Render(lipgloss.JoinVertical(lipgloss.Left, rows...))
}

func (sd StepDefinition) ViewValue() string {
	if sd.Step.Filled() {
		return "configured"
	}

	return menuValueStyle.Copy().Foreground(lipgloss.Color("240")).Render("not configured")
}
