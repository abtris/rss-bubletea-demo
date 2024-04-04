package main

import (
	"fmt"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	file, _ := os.Open("data/podcast.xml")
	defer file.Close()
	fp := gofeed.NewParser()
	feed, _ := fp.Parse(file)
	fmt.Println(feed.Title)
	for _, episode := range feed.Items {
		fmt.Println(episode.Title)
	}
}
