package message

import (
	"bubbletea-reducer/app/model"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// Exit is a message to close the app
type Exit[M model.Exit] int

func (_ Exit[M]) Reduce(m *M, msg tea.Msg) (*M, tea.Cmd) {
	data := any(m).(*model.Exit)
	switch count := msg.(Exit[M]); {
	case count >= 0:
		data.Count = int(count)
		return any(data).(*M), func() tea.Msg {
			time.Sleep(time.Second / 2)
			return Exit[M](count - 1)
		}
	default:
		return any(data).(*M), tea.Quit
	}
}
