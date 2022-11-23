package app

import (
	"bubbletea-reducer/app/model"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	m.Model = &model.Model{
		Text: "",
		Exit: &model.Exit{Count: -1},
		Key:  &model.Key{TeaCount: 0},
	}

	return func() tea.Msg {
		// Send a keypress
		return tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune{'w'},
			Alt:   false,
		}
	}
}
