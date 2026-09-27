package main

import (
	"net/url"

	"sync"
)

func configure(baseURL string, maxConcurrency int) (*config, error) {
	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err

	}
	configured := config{
		pages:              map[string]PageData{},
		baseURL:            parsedBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}
	return &configured, nil
}
