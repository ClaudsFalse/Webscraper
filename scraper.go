package main

import (
	"fmt"
	"os"

	// import Colly
	"github.com/gocolly/colly"
)

type Recipe struct {
	Url, Name string
}

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(os.Getenv("URL"))
}
