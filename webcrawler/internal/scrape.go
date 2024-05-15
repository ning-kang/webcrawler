package internal

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

func ScrapeBooks(url string) {
	fmt.Println("Scraping books from", url)
	html, err := get(url)
	if err != nil {
		log.Fatal("Unable to scrape book:", err)
	}
	traverseHTML(html)
}

func get(url string) (*html.Node, error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	html, err := html.Parse(res.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}

func traverseHTML(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseHTML(c)
	}
}
