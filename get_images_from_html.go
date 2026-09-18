package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error trying to read from html: %w", err)
	}
	var imgURLs []string

	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists {
			fmt.Println("error trying to get src attribute from img tag")
			return // exit from this call and continues to next
		}

		if strings.TrimSpace(src) == "" {
			return // exit from this call and continues to next
		}

		parsedLink, err := url.Parse(src)
		if err != nil {
			return // exit from this call and continues to next
		}
		absolutePath := baseURL.ResolveReference(parsedLink)
		imgURLs = append(imgURLs, absolutePath.String())

	})

	return imgURLs, nil
}
