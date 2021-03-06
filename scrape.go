package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func scrape(url string) string {
	defer timeTrack(time.Now(), "factorial - scrape")
	// Request the HTML page.
	var obj string
	res, err := http.Get(url)
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
	defer timeTrack(time.Now(), "factorial - scanner")
	var value string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a topic: ")
	for scanner.Scan() {
		value = string(scanner.Text())
		break
	}
	return "https://flipboard.com/topic/" + value
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
