package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    string
		expected    string
		hasError    bool
		errContains string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
			hasError: false,
		},
		{
			name:     "remove scheme and slash",
			inputURL: "https://www.boot.dev/blog/path/",
			expected: "www.boot.dev/blog/path",
			hasError: false,
		},
		{
			name:     "remove Upper Case",
			inputURL: "http://WWW.BOOT.DEV/blog/path",
			expected: "www.boot.dev/blog/path",
			hasError: false,
		},
		{
			name:     "remove Upper Case with slash",
			inputURL: "http://WWW.BOOT.DEV/blog/path/",
			expected: "www.boot.dev/blog/path",
			hasError: false,
		},
		{
			name:        "Invalid Url",
			inputURL:    "://invalid-url",
			expected:    "",
			hasError:    true,
			errContains: "Error trying to parse url",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if tc.hasError {
				if err == nil {
					t.Errorf("Test %v - '%s' FAIL: expected error: got nil", i, tc.name)
					return
				}

				if tc.errContains != "" && !strings.Contains(err.Error(), tc.errContains) {
					t.Errorf("Test %v - '%s' FAIL: expected error: %v, got error: %v", i, tc.name, tc.errContains, err)
				}
				return
			}
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
