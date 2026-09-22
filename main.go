package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Printf("no website provided")
		os.Exit(1)
	} else if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl := argsWithoutProg[0]
	pages := make(map[string]int)
	crawlPage(baseUrl, baseUrl, pages)

	for k, v := range pages {
		fmt.Printf("Visited this page %s for %d/n", k, v)
	}
}
