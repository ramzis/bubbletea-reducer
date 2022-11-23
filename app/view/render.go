package view

import (
	"bubbletea-reducer/app/model"
	"strings"
)

func Render(m *model.Model) string {
	if m.Count < 0 {
		m.Text = "Hello, Bubble Tea!" + strings.Repeat("🧋", m.TeaCount)
	} else if m.Count > 0 {
		m.Text = "Exiting" + strings.Repeat(".", int(m.Count))
	} else {
		m.Text = "Bye!"
	}

	return m.Text
}
