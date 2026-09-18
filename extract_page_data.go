package main

import (
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	heading := getHeadingFromHTML(html)
	firstParagraph := getFirstParagraphFromHTML(html)

	// if parsing fail then bail early with only fields that rely on raw html
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		return PageData{
			URL:            pageURL,
			Heading:        heading,
			FirstParagraph: firstParagraph,
			OutgoingLinks:  nil,
			ImageURLs:      nil,
		}
	}

	outgoingLinks, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		outgoingLinks = nil
	}

	imageURLs, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		imageURLs = nil
	}

	return PageData{
		URL:            pageURL,
		Heading:        heading,
		FirstParagraph: firstParagraph,
		OutgoingLinks:  outgoingLinks,
		ImageURLs:      imageURLs,
	}
}
