package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	var urls []string
	var links []string
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error trying to read from html - %w", err)
	}

	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			links = append(links, href)
		}
	})

	for _, link := range links {
		parsedURL, err := url.Parse(link)
		if err != nil {
			return nil, fmt.Errorf("Error parsing href: %w", err)
		}
		resolvedURL := baseURL.ResolveReference(parsedURL)
		urls = append(urls, resolvedURL.String())
	}

	return urls, nil

}
