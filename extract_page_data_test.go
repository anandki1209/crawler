package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  PageData
	}{
		{
			name:     "with All Attributes",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
        					<h1>Test Title</h1>
            				<p>This is the first paragraph.</p>
                			<a href="/link1">Link 1</a>
                   			<img src="/image1.jpg" alt="Image 1">
                     	</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://crawler-test.com/link1"},
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
		{
			name:     "with  missing Paragraph",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
        					<h1>Test Title</h1>

                			<a href="/link1">Link 1</a>
                   			<img src="/image1.jpg" alt="Image 1">
                     	</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "",
				OutgoingLinks:  []string{"https://crawler-test.com/link1"},
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
		{
			name:     "with  missing anchor tag",
			inputURL: "https://crawler-test.com",
			inputBody: `<html><body>
        					<h1>Test Title</h1>
             				<p>This is the first paragraph.</p>
                   			<img src="/image1.jpg" alt="Image 1">
                     	</body></html>`,
			expected: PageData{
				URL:            "https://crawler-test.com",
				Heading:        "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  nil,
				ImageURLs:      []string{"https://crawler-test.com/image1.jpg"},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputBody, tc.inputURL)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, actual)
			}
		})
	}

}
