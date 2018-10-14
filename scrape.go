package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://flipboard.com/topic/cryptocurrency")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if x, _ := s.Attr("type"); x == "application/ld+json" {
			fmt.Printf(s.Text()) // get
		}

	})

}

func main() {
	ExampleScrape()
}
