package webpage

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// Get gets text content from a URL
func Get(url string, selector string) (textContent string, err error) {
	// Get URL
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}

	// Find the first selector's text content
	return doc.Find(selector).First().Text(), nil
}
