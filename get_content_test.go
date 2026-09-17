package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{
			name:      "extracting heading 1",
			inputBody: "<html><body><h1>Test Title</h1></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "extracting heading 2",
			inputBody: "<html><body><h2>Test Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "extracting form heading1 and heading2 ",
			inputBody: "<html><body><h1>Test Title</h1><h2>Side Title</h2></body></html>",
			expected:  "Test Title",
		},
		{
			name:      "extracting from heading 3",
			inputBody: "<html><body><h3>Test Title</h3></body></html>",
			expected:  "",
		},
		{
			name:      "extracting from no heading",
			inputBody: "<html><body><p>Test Title</p></body></html>",
			expected:  "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getHeadingFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test case - %v fail, expected %q, got %q", i, tc.expected, actual)
			}
		})
	}

}

func TestGetFirstParagraphFromHTMLMainPriority(t *testing.T) {
	tests := []struct {
		name      string
		inputBody string
		expected  string
	}{
		{name: "extracting main paragraph",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<p>Main paragraph.</p>
				</main>
			</body></html>`,

			expected: "Main paragraph.",
		},
		{name: "extracting outside paragraph",
			inputBody: `<html><body>
				<p>Outside paragraph.</p>
				<main>
					<h1>Main paragraph.</h1>
				</main>
			</body></html>`,

			expected: "Outside paragraph.",
		},
		{name: "extracting no paragraph",
			inputBody: `<html><body>
				<title>Outside paragraph.</title>
				<main>
					<h1>Main paragraph.</h1>
				</main>
			</body></html>`,

			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(tc.inputBody)
			if actual != tc.expected {
				t.Errorf("Test case- %v failed, expected %q, got %q", i, tc.expected, actual)
				return
			}

		})
	}

}
