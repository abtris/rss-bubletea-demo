package tui

import (
	"log"

	tea "charm.land/bubbletea/v2"
	"charm.land/glamour/v2"
)

func (m model) View() tea.View {
	v := tea.NewView(m.render())
	v.AltScreen = true
	v.MouseMode = tea.MouseModeCellMotion
	return v
}

func (m model) render() string {
	var s string
	if len(m.choice) > 0 {
		s += "# " + title
		s += "\n## " + m.choice
		s += "\n\n"
		s += m.content
		renderer, err := glamour.NewTermRenderer(
			glamour.WithEnvironmentConfig(),
			glamour.WithWordWrap(width),
		)
		if err != nil {
			return ""
		}
		out, err := renderer.Render(s)
		if err != nil {
			log.Println(err)
		}
		return out
	}
	return m.list.View()
}
