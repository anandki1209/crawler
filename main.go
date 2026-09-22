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
	fmt.Printf("starting crawl of: %s", baseUrl)

}
