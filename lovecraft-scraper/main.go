package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

/*
The html on hplovecraft.com is truly a sight to behold. I can't
help but wonder if it was written to torment the souls of web developers
who might one day look upon the page source and scream as they realise
that what has been created here is truly an abomination.

Seriously. Go have a look. It'll destroy your mind.
*/

var (
	url = "http://www.hplovecraft.com/writings/texts/"
	out = fmt.Sprintf("%s/tmp/lovecraft.txt", os.Getenv("HOME"))
)

func main() {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Traverse the index and pull out all the links in the fiction section
	doc.Find("li a").Each(func(i int, s *goquery.Selection) {
		h, _ := s.Attr("href")
		if strings.Contains(h, "fiction/") {
			fmt.Println("Scraping ", h)
			article, err := goquery.NewDocument(fmt.Sprintf("%s%s", url, h))
			if err != nil {
				log.Println(err)
			}

			/*
				the "done" value is necessary because every page contains 2 divs
				that open with the attr align set to justify. One just before the
				other. This looks like bad template writing to me.
			*/
			done := false
			article.Find("div").Each(func(i int, d *goquery.Selection) {
				v, ok := d.Attr("align")
				if ok && v == "justify" && !done {
					fmt.Fprint(f, d.Text())
					done = true
				}
			})
		}
	})
	fmt.Println("Don't forget to copy ~/tmp/lovecraft.txt to the trumpcraft sources folder.")
}
