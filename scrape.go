package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type information struct {
	id       string `json:"id"`
	title    string `json:"title"`
	location string `json:"location"`
	summary  string `json:"summary"`
}

func createFile(file []byte) {
	this := ioutil.WriteFile("output.json", file, 0644)
	if err := this; err != nil {
		panic(err)
	}
}

func serializeJSON(info []information) {
	fmt.Println("Serialization")
	infoJSON, _ := json.Marshal(info)
	createFile(infoJSON)
	fmt.Println("Completed")
	fmt.Println(string(infoJSON))
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	infos := []information{}

	// On every a element which has href attribute call callback
	c.OnHTML("body > div > div.row.header-box", func(e *colly.HTMLElement) {
		//Child texts
		info := information{}
		info.id = e.ChildText("id")
		info.title = e.ChildText("title")
		info.location = e.ChildText("location")
		info.summary = e.ChildText("summary")

		// Print link
		fmt.Printf("Job listing: %q\n", info.id, info.title, info.location, info.summary)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.indeed.com/")
	serializeJSON(infos)
}
