package main

import (
	"fmt"
	"log"
	"strings"
	
	"github.com/gocolly/colly/v2"
)

func main() {
	collector:=colly.NewCollector()

	url:="https://nasional.kompas.com/"

	collector.OnHTML(".articleItem", func(h *colly.HTMLElement) {
		title:=h.ChildText("h2.articleTitle")
		date:=h.ChildText("div.articlePost-date")
		link:=h.ChildAttr("a", "href")
		link=strings.TrimSpace(link)
		date=strings.TrimSpace(date)
		title=strings.TrimSpace(title)
		
		fmt.Printf("Article title : %s\nDate : %s\nHere's the link to the article :\n%s\n\n", title, date, link)
	})

	err:=collector.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}