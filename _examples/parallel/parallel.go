package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// Limit the maximum parallelism to 5
	// This is necessary if the goroutines are dynamically
	// created to control the limit of simultaneous requests.
	//
	// Parallelism can be controlled also by spawning fixed
	// number of go routines.
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	// MaxDepth is 2, so only the links on the scraped page
	// and links on those pages are visited
	c.MaxDepth = 2

	// On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	// link := e.Attr("href")
	// 	link := e.Attr("h4[span]")
	// 	// Print link
	// 	fmt.Println(link)
	// 	// Visit link found on page on a new thread
	// 	// go e.Request.Visit(link)
	// })
	c.OnHTML(".card-group", func(elem *colly.HTMLElement) {
		// div.creator-names > span
		// temp := e.ChildText("span.aria-count-up")
		// divClass := "card-body card-block"
		nestedTestData := elem.ChildText("div")
		fmt.Println("xx:", nestedTestData)
		doc, _ := goquery.NewDocumentFromReader(bytes.NewBuffer([]byte(nestedTestData)))
		e := &colly.HTMLElement{
			DOM: doc.First(),
		}
		type nested struct {
			String string  `selector:"div > span"`
			Struct *nested `selector:"div > div"`
		}
		s := nested{}
		if err := e.Unmarshal(&s); err != nil {
			fmt.Println("Cannot unmarshal struct: " + err.Error())
		}
		// if s.String != "a" {
		// 	t.Errorf(`Invalid data for String: %q, expected "a"`, s.String)
		// }
		fmt.Println(s.String)
		fmt.Println(s.Struct.Struct.String)
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("https://www.etherchain.org/")
	// Wait until threads are finished
	c.Wait()
}
