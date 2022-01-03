package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func createCrawler(domains []string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
		colly.UserAgent("golatx_bot v0.0.1"),
	)

	return c
}

func crawl(c *colly.Collector, link string) int {
	count := 0

	c.OnRequest(func(r *colly.Request) {
		count++
		fmt.Println("Scrawling:	", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		name := strings.Split(r.FileName(), ".")[0]
		host := r.Request.URL.Hostname()
		saveToDisk(name, string(r.Body), host)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if e.Attr("class") == "question-hyperlink" {
			c.Visit(e.Request.AbsoluteURL(link))
		}
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error:", r.StatusCode, err)
	})

	c.Visit(link)

	return count
}
