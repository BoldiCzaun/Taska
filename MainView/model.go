package mainview

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func InitialModel() Model {
	return Model{
		choices: []string{"Item 1", "Item 2", "Item 3"},

		selected: make(map[int]struct{}),
	}
}
func (m Model) Init() tea.Cmd {
	return nil
}
