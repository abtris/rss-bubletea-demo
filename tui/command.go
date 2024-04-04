package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct {
	err error
}

type scanMsg struct {
	items []list.Item
}

func (m model) scanCmd() tea.Cmd {
	return func() tea.Msg {
		return scanMsg{items: []list.Item{}}
	}
}
