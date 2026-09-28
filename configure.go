package main

import (
	"net/url"

	"sync"
)

type config struct {
	pages              map[string]PageData // keep track of pages we have crwaled
	baseURL            *url.URL            // keep track of original base baseURL
	mu                 *sync.Mutex         // ensures pages map is thread safe
	concurrencyControl chan struct{}       // buffered channel of  empty struct
	wg                 *sync.WaitGroup     // ensures main wait for all goroutines to finish first
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true

}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data

}

// immediatly unlock mutex since other recursive goroutines need it
func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func configure(baseURL string, maxConcurrency, maxPagesLimit int) (*config, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err

	}
	return &config{
		pages:              map[string]PageData{},
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPagesLimit,
	}, nil

}
