package internal

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
)

func ScrapePage(url string) {
	log.Println("Scraping page:", url)
	res := get(url)

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Status code error:", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to get html document:", err)
	}

	var list []string

	doc.Find(".product_pod").Each(func(i int, s *goquery.Selection) {
		bookUrl, ok := s.Find("a").Attr("href")
		if ok {
			list = append(list, "https://books.toscrape.com/"+bookUrl)
		}
	})

	var books []Book

	for _, l := range list {
		books = scrapeBook(books, l)
	}

	fmt.Print(books)
}

func scrapeBook(b []Book, url string) []Book {
	log.Println("Scraping book:", url)
	res := get(url)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal("Status code error:", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Failed to get html document:", err)
	}

	book := Book{}

	p := doc.Find(".product_page")
	m := p.Find(".product_main")

	book.Title = m.Find("h1").Text()

	priceStr := m.Find(".price_color").Text()

	price, err := strconv.ParseFloat(trimFirstRune(priceStr), 32)
	if err != nil {
		log.Fatal("Not able to parse price:", priceStr)
	}
	book.Price = float32(price)

	book.Description = p.Find("div#product_description").Next().Text()

	p.Find("table").Each(func(i int, table *goquery.Selection) {
		table.Find("tr").Each(func(i int, row *goquery.Selection) {
			th := row.Find("th").Text()
			switch th {
			case "UPC":
				book.UPC = row.Find("td").Text()
			case "Tax":
				taxStr := row.Find("td").Text()
				tax, err := strconv.ParseFloat(trimFirstRune(taxStr), 32)
				if err != nil {
					log.Fatal("Not able to parse tax:", taxStr)
				}
				book.Tax = float32(tax)
			case "Availability":
				book.Availability = row.Find("td").Text()
			}
		})
	})

	return append(b, book)
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
