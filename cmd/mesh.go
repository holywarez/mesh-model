package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"mesh-model/internal/tui"
	"os"
)

func main() {
	p := tea.NewProgram(tui.New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
