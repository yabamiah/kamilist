package internal

import "github.com/charmbracelet/lipgloss"

var msg = lipgloss.NewStyle().
	SetString(" Welcome to かみlist! ").
	Background(lipgloss.Color("#D90429")).
	Bold(true).
	Margin(2, 4).
	String()

func (m *Model) View() string {
	return msg
}