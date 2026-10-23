package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abtris/rss-bubbletea/tui"
	tea "charm.land/bubbletea/v2"
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

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		log.Fatal("start failed: ", err)
	}
}
