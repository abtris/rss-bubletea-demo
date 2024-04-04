package tui

import (
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
	for _, rssItem := range data.Items {
		i := item{
			title: rssItem.Title,
			desc:  rssItem.Description,
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
