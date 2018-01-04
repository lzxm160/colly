package main

import (
	"fmt"

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
	c.OnHTML("h4[class=card-title]", func(e *colly.HTMLElement) {
		// if e.DOM.Find("section.course-info").Length() == 0 {
		// 	return
		// }
		// title := strings.Split(e.ChildText(".course-title"), "\n")[0]
		// course_id := e.ChildAttr("input[name=course_id]", "value")
		// start_date, _ := time.Parse(DATE_FORMAT, e.ChildText("span.start-date"))
		// end_date, _ := time.Parse(DATE_FORMAT, e.ChildText("span.final-date"))
		temp := e.ChildText("span.aria-count-up")
		fmt.Println("xx:", temp)
		// var run string
		// if len(strings.Split(course_id, "_")) > 1 {
		// 	run = strings.Split(course_id, "_")[1]
		// }
		// course := Course{
		// 	CourseID:  course_id,
		// 	Run:       run,
		// 	Name:      title,
		// 	Number:    e.ChildText("span.course-number"),
		// 	StartDate: &start_date,
		// 	EndDate:   &end_date,
		// 	URL:       fmt.Sprintf("/courses/%s/about", course_id),
		// }
		// courses = append(courses, course)
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit("https://www.etherchain.org/")
	// Wait until threads are finished
	c.Wait()
}
