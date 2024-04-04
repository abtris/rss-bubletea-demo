package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abtris/rss-bubbletea/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mmcdole/gofeed"
)

func main() {
	file, _ := os.Open("data/podcast.xml")
	defer file.Close()
	fp := gofeed.NewParser()
	feed, err := fp.Parse(file)
	if err != nil {
		log.Fatal("parse feed failed", err)
	}

	model, err := tui.NewModel(feed)
	if err != nil {
		log.Fatal("create model failed", err)
	}

	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	p := tea.NewProgram(model, tea.WithAltScreen(), tea.WithMouseCellMotion())
	if err := p.Start(); err != nil {
		log.Fatal("start failed: ", err)
	}
}
