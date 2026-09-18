package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	heading := doc.Find("h1, h2").First().Text()
	return strings.TrimSpace(heading)

}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	main := doc.Find("main")
	var para string
	para = main.Find("p").First().Text()
	if para == "" {
		para = doc.Find("p").First().Text()
	}
	return strings.TrimSpace(para)

}
