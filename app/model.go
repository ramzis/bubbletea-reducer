package app

import (
	"bubbletea-reducer/app/model"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	*model.Model
}

func Build() tea.Model {
	return &Model{}
}
