package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 3 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := argsWithoutProg[0]
	maxConcurrency, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg, err := configure(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// blocking main until all goroutines gets completed
	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	for k := range cfg.pages {
		fmt.Printf("found: %s\n", k)
	}

}
