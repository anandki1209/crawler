package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "BootCrawler/1.0")

	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close() // use defer to close res.Body whenever you access res any attribute
	if res.StatusCode >= 400 {
		return "", fmt.Errorf("error in trying to crawl web page")
	}

	if contentHeader := res.Header.Get("content-type"); !strings.HasPrefix(contentHeader, "text/html") {
		return "", fmt.Errorf("content-type haeder not present")
	}

	htmlResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(htmlResponse), nil

}
