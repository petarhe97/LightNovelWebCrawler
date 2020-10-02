package esjscraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

const (
	esjzoneURL = "https://www.esjzone.cc/list/"
)

// EsjScraper exported
type EsjScraper struct {
	target string
	Titles []string
}

// NewEsjScraper : EsjScraper constructor exported
func NewEsjScraper() *EsjScraper {
	es := EsjScraper{target: esjzoneURL}
	return &es
}

func (scraper *EsjScraper) addToList(element string) {
	scraper.Titles = append(scraper.Titles, element)
}

// Scrape : scrapes the data of esjzone
func (scraper *EsjScraper) Scrape() {
	c := colly.NewCollector()

	c.OnHTML(".card", func(e *colly.HTMLElement) {
		scraper.addToList(e.Attr("title"))
	})

	c.Visit("https://www.esjzone.cc/list/")
	fmt.Println(len(scraper.Titles))

	for i := 0; i < len(scraper.Titles); i++ {
		fmt.Println(scraper.Titles[i])
	}
}
