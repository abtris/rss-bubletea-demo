package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.detail = true
				m.choice = string(i.title)
				m.content = string(i.desc)
			}
			return m, nil
		case "b":
			if m.detail {
				m.choice = ""
				m.content = ""
			}
		case "p":
			if m.detail {
				changeIndex := m.list.Index() + 1
				if changeIndex <= 0 {
					changeIndex = 0
				}
				m.list.Select(changeIndex)
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.choice = string(i.title)
					m.content = string(i.desc)
				}
			}
		case "n":
			if m.detail {
				changeIndex := m.list.Index() - 1
				maxLength := len(m.list.Items())
				if changeIndex > (maxLength - 1) {
					changeIndex = maxLength - 1
				}
				m.list.Select(changeIndex)
				i, ok := m.list.SelectedItem().(item)
				if ok {
					m.choice = string(i.title)
					m.content = string(i.desc)
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
