package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("Error trying to parse url: %w", err)
	}

	fullPath := parsedUrl.Host + parsedUrl.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
