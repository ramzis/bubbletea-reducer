package message

import (
	"bubbletea-reducer/app/model"
	tea "github.com/charmbracelet/bubbletea"
	"math"
)

// Key is a message to wrap the default tea.KeyMsg
type Key[D model.Key] tea.KeyMsg

func (_ Key[M]) Reduce(m *M, msg tea.Msg) (*M, tea.Cmd) {
	data := any(m).(*model.Key)
	switch tea.KeyMsg(msg.(Key[M])).String() {
	case "q", "esc", "ctrl+c":
		return any(data).(*M), func() tea.Msg {
			return Exit[model.Exit](3)
		}
	case "w":
		data.TeaCount += 1
		return any(data).(*M), nil
	case "s":
		data.TeaCount = int(math.Max(0, float64(data.TeaCount-1)))
		return any(data).(*M), nil
	default:
		return any(data).(*M), nil
	}
}
