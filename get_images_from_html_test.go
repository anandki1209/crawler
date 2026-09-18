package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:      "relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:      "absolute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="https://crawler-test.com/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		{
			name:     "multiple links",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body><img src="https://crawler-test.com/logo.png" alt="Logo">
			<img src="/mascot.png" alt="mascot"></body></html>`,
			expected: []string{"https://crawler-test.com/logo.png", "https://crawler-test.com/mascot.png"},
		},
		{
			name:      "missing attributes",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img alt="Logo"></body></html>`,
			expected:  nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}
			actual, err := getImagesFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}

		})
	}

}
