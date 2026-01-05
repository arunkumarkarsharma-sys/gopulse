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
		fmt.Println(" Error:", err)
		return
	}

	fmt.Println(" Checking websites concurrently")

	results := checker.CheckURLsConcurrently(urls)

	for i := 0; i < len(urls); i++ {
		result := <-results

		if result.Error != nil {
			fmt.Printf(" %s → %s (%v)\n",
				result.URL,
				result.Status,
				result.Error,
			)
		} else {
			fmt.Printf(" %s → %s (%v)\n",
				result.URL,
				result.Status,
				result.ResponseTime,
			)
		}
	}
}
