package main

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func Start(link string) int {
	pol := genPolicy(link)
	return crawl(createCrawler(pol.Domains), link, &pol)
}

func createCrawler(domains []string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(domains...),
		colly.UserAgent("golatx_bot v0.0.1"),
	)

	return c
}

func genPolicy(link string) Policy {
	domain, err := url.Parse(link)
	checkError(err)

	policy := readJson(filepath.Join("./", "data", "policies.json"))
	return policy[domain.Hostname()]

	/**
	 *Receive the json and parse it, then convert it into a struct
	 *Return the struct with each of the fields mapped accordingly
	 *This function should run before crawl func since latter depends on former
	**/
}

func crawl(c *colly.Collector, link string, pol *Policy) int {
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

		if len(pol.LinkClasses) == 0 {
			c.Visit(e.Request.AbsoluteURL(link))
		} else {
			for _, class := range pol.LinkClasses {
				if e.Attr("class") == class {
					c.Visit(e.Request.AbsoluteURL(link))
				}
			}
		}

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.StatusCode, err)
	})

	c.Limit(&colly.LimitRule{
		RandomDelay: 5 * time.Second,
	})

	c.Visit(link)

	return count
}
