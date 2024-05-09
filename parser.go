package main

import (
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://kbd.news/rss2.php")
}
