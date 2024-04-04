package tui

func (m model) View() string {
	var s string
	if len(m.choice) > 0 {
		s += m.choice
		s += "\n\n"
		s += m.content
		return s
	}
	return m.list.View()
}
