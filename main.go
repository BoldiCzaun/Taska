package main

import (
	"fmt"
	"os"

	mainview "github.com/BoldiCzaun/Taskmanager/MainView"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(mainview.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
