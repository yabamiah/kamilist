package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yabamiah/kamilist/internal"
)

func main() {
	m := internal.New()
	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
