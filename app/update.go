package app

import (
	"bubbletea-reducer/app/message"
	"bubbletea-reducer/app/model"
	"bubbletea-reducer/pkg/reducer"
	tea "github.com/charmbracelet/bubbletea"
)

// Update performs partial Model updates based on the tea.Msg type
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmd := tea.Cmd(nil)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.Key, cmd = reducer.Reduce(m.Key, message.Key[model.Key](msg))
	case message.Exit[model.Exit]:
		m.Exit, cmd = reducer.Reduce(m.Exit, msg)
	}
	return m, cmd
}
