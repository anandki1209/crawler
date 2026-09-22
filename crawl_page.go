package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	if parsedCurrentURL.Host != parsedBaseURL.Host {
		return
	}

	normalizedRawCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, ok := pages[normalizedRawCurrentURL]
	if ok {
		pages[normalizedRawCurrentURL]++
		return
	}
	pages[normalizedRawCurrentURL] = 1

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("crawling this page : %s/n", rawCurrentURL)

	urls, err := getURLsFromHTML(html, parsedBaseURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}

}
