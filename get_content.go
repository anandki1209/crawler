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

	heading := doc.Find("h1")
	if heading.Length() == 0 {
		return doc.Find("h2").First().Text()
	}
	return heading.First().Text()

}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}

	para := doc.Find("main p")
	if para.Length() == 0 {
		return doc.Find("p").First().Text()
	}
	return para.First().Text()
}
