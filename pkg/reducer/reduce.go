package reducer

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type Reducer[M any] interface {
	Reduce(m *M, msg tea.Msg) (*M, tea.Cmd)
}

// Reduce resolves and applies a Model reducer for a given tea.Msg
func Reduce[M any](m *M, msg tea.Msg) (*M, tea.Cmd) {
	if reducer, ok := msg.(Reducer[M]); !ok {
		panic(fmt.Sprintf("reducer not found for %v", m))
	} else {
		return reducer.Reduce(m, msg)
	}
}
