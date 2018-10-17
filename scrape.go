package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func scrape(tp string) string {
	// Request the HTML page.
	var obj string
	res, err := http.Get("https://flipboard.com/topic/" + tp)
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
		/*
			Loop through script tag and find "type" attribute that
			is "application/ld+json" because json data is in here
		*/
		if x, _ := s.Attr("type"); x == "application/ld+json" {
			obj = s.Text() // get
		}

	})

	return obj

}

func scanner() string {
	var value string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a topic: ")
	for scanner.Scan() {
		value = string(scanner.Text())
		break
	}
	return value
}

func main() {
	s := scanner()
	fmt.Println("Searching...")
	fmt.Println(scrape(s))
}

/*
TODO:
 - convert string into json in struct
 - run a subroutine to access news url
from proxy urls obtain from json
*/
