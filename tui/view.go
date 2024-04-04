package tui

import (
	"log"

	"github.com/charmbracelet/glamour"
)

func (m model) View() string {
	var s string
	if len(m.choice) > 0 {
		s += "# " + title
		s += "\n## " + m.choice
		s += "\n\n"
		s += m.content
		renderer, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
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
