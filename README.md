# Crawler

A small concurrent web crawler written in Go. Starting from a URL, it visits pages on the same host, extracts page content and links, and saves the results as JSON.

## Requirements

- Go 1.26.3 or later

## Run

Pass the starting URL, maximum number of concurrent requests, and maximum number of pages to crawl:

```sh
go run . https://example.com 5 100
```

The arguments are positional:

1. Starting URL (include `https://` or `http://`)
2. Maximum concurrency
3. Maximum pages

You can also build an executable and run it:

```sh
go build -o crawler .
./crawler https://example.com 5 100
```

The crawler prints visited URLs and writes `report.json` in the current directory when it finishes. Each entry in the JSON array contains:

- `url`: page URL
- `heading`: first `h1` or `h2` heading found
- `first_paragraph`: first paragraph in `<main>`, or the first paragraph on the page if there is no paragraph in `<main>`
- `outgoing_links`: links found on the page, resolved to absolute URLs
- `image_urls`: image `src` values, resolved to absolute URLs

The crawler only follows links whose host matches the starting URL's host. The page limit applies to discovered pages; the concurrency argument controls how many crawl goroutines can run at once.

## Tests

Run the test suite with:

```sh
go test ./...
```
