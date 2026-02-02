package main

import (
	"fmt"
	"gopulse/checker"
	"gopulse/input"
)

func main() {
	fmt.Println(" GoPulse CLI started")

	urls, err := input.ReadURLsFromFile("urls.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Checking websites concurrently")

	results := checker.CheckURLsConcurrently(urls)

	for _, result := range results {
		if result.Error != nil {
			fmt.Printf("%s -> ERROR: %v\n", result.Url, result.Error)
		} else {
			fmt.Printf("%s -> %s\n", result.Url, result.Status)
		}
	}
}
