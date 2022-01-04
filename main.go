package main

import (
	"fmt"
	"os"

	"golatx/crawler"
)

func main() {

	switch os.Args[1] {
	case "-i":
		crawlOne(os.Args[2])
	case "-f":
		crawlFromFile(os.Args[2])
	default:
		panic("Invalid arguments passed")
	}

}

func crawlOne(link string) {
	fmt.Println("Starting...")

	crawler.Start(link)
}

func crawlFromFile(url string) {
	fmt.Println("Starting...")
}

/*
* Option to parse args passed via terminal when running
* Alternatively pass a json/text file containing all urls to crawl and their options
* For each of the urls, implement a goroutine for crawling
* Implement channels and select options
 */
