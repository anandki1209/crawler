package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {

	var urls []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error trying to read from html - %w", err)
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			fmt.Println("error trying to get href attribute from anchor tag")
			return // exit form this callback iteration and move to next
		}

		if strings.TrimSpace(href) == "" {
			return // exit form this callback iteration and move to next
		}

		parsedURL, err := url.Parse(href)
		if err != nil {
			return // ends this callback and continues to next iteration
		}
		resolvedURL := baseURL.ResolveReference(parsedURL)
		urls = append(urls, resolvedURL.String())

	})

	return urls, nil

}
