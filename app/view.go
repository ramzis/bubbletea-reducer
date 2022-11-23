package app

import (
	"bubbletea-reducer/app/view"
)

func (m *Model) View() string {
	return view.Render(m.Model)
}
