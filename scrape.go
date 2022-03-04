package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gocolly/colly"
)

type Jobs struct {
	title    string
	location string
	summary  string
}

func createFile(file []Jobs) {
	jsonFile, _ := json.MarshalIndent(file, "", " ")
	_ = ioutil.WriteFile("Jobs.json", jsonFile, 0644)
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	c.SetRequestTimeout(60 * time.Second)
	page := make([]Jobs, 0)

	// On every a element which has href attribute call callback
	c.OnHTML("table.article-table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, e *colly.HTMLElement) {
			info := Jobs{}
			info.title = e.ChildText("body")
			info.location = e.ChildText("td")
			info.summary = e.ChildText("td")
			page = append(page, info)
		})
	})

	// Before making a request print "Visiting ..."

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	c.Visit("https://www.indeed.com/")

	createFile(page)
}
