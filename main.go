package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	fmt.Println("Hello, world.")
	c := colly.NewCollector()
	var lst []string

	c.OnHTML(".card", func(e *colly.HTMLElement) {
		//e.Request.Visit(e.Attr("href"))
		lst = append(lst, e.Attr("title"))
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })

	c.Visit("https://www.esjzone.cc/list/")

	for i := 0; i < len(lst); i++ {
		fmt.Println(lst[i])
	}
}
