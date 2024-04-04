package tui

import (
	"log"

	"github.com/charmbracelet/glamour"
)

func (m model) View() string {
	var s string
	if len(m.choice) > 0 {
		s += m.choice
		s += "\n\n"
		s += m.content
		out, err := glamour.Render(s, "dark")
		if err != nil {
			log.Println(err)
		}
		return out
	}
	return m.list.View()
}
