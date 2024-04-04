package tui

import (
	"log"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/bubbles/list"
	"github.com/mmcdole/gofeed"
)

type model struct {
	list     list.Model
	choice   string
	content  string
	detail   bool
	quitting bool
}

const width = 80
const height = 40
const title = "RSS Reader"

func NewModel(data *gofeed.Feed) (*model, error) {
	var items []list.Item
	converter := md.NewConverter("", true, nil)
	for _, rssItem := range data.Items {
		markdown, err := converter.ConvertString(rssItem.Description)
		if err != nil {
			log.Println("Convert to markdown", err)
		}
		i := item{
			title: rssItem.Title,
			desc:  "Published at " + rssItem.Published + "\n\n" + markdown,
		}
		items = append(items, i)
	}
	l := list.New(items, list.NewDefaultDelegate(), width, height)
	l.Title = title

	return &model{
		list:    l,
		choice:  "",
		content: "",
		detail:  false,
	}, nil
}
