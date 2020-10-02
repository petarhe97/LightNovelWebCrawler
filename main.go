package main

import (
	"fmt"

	"github.com/lightnovelwebcrawler/esjscraper"
)

func main() {
	scraper := esjscraper.NewEsjScraper()
	scraper.Scrape()
	for i := 0; i < len(scraper.Titles); i++ {
		fmt.Println(scraper.Titles[i])
	}
}
