package tui

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/bubbles/list"
	"github.com/mmcdole/gofeed"
)

type state int

type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func NewModel(data *gofeed.Feed) (*model, error) {
	var items []list.Item
	converter := md.NewConverter("", true, nil)
	for _, rssItem := range data.Items {
		markdown, err := converter.ConvertString(rssItem.Description)
		if err != nil {
			log.Printf("Convert to markdown", err)
		}
		i := item{
			title: rssItem.Title,
			desc:  markdown,
		}
		items = append(items, i)
	}
	width, height := 80, 40
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = "RSS Reader"

	return &model{
		list: l,
	}, nil
}
