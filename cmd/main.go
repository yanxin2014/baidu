package main

import (
	"log"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"os"
	"fmt"
)

const SEPARATOR = "\u0001"


type Item struct{
	id string
	href string
	name string
	date string
}

func ExampleScrape() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var item Item
	var id string = "-1"
	// Request the HTML page.
	for {
		if len(id) <= 0 {
			break
		}
		res, err := http.Get("http://wjsou.com/sindex.php?id=" + id)
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

		// Find the review items
		id = ""
		doc.Find("h4").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			a := s.Find(" a.fname")
			id, _ = a.Attr("id")
			href, _ := a.Attr("href")
			item.id = id
			item.href = href
			item.name = a.Text();
			date := a.Next().Text();
			item.date = date;
			file.WriteString(item.id + SEPARATOR + item.name + SEPARATOR + item.href + SEPARATOR + item.date + "\r\n")
		})
		fmt.Println("id : " + id)
	}
}

func init() {
}

func main() {
	ExampleScrape()

}