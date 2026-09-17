package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		hasError  bool
	}{
		{name: "extracting a url from html",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		{name: "extracting multiple url from html",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
			<a href="https://crawler-test.com"><span>Boot.dev
			</span>
			<a href="https://boot.dev"><span>Boot.dev</span>
			</a></body></html>`,
			expected: []string{"https://crawler-test.com", "https://boot.dev"},
		},
		{name: "extracting no url from html",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><h1>Blog</h1</body></html>`,
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

			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
