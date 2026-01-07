package checker

import (
	"net/http"
	"sync"
	"time"
)

// CheckResult holds the result of a URL check

type CheckResult struct {
	URL          string
	Status       string
	ResponseTime time.Duration
	Error        error
}

// reusable HTTP client

var client = http.Client{
	Timeout: 10 * time.Second,
}

// CheckURL checks a single URL

func CheckURL(url string) CheckResult {
	start := time.Now()

	resp, err := client.Get(url)
	if err != nil {
		return CheckResult{
			URL:    url,
			Status: "DOWN",
			Error:  err,
		}
	}
	defer resp.Body.Close()

	return CheckResult{
		URL:          url,
		Status:       resp.Status,
		ResponseTime: time.Since(start),
		Error:        nil,
	}
}

// CheckURLsConcurrently checks multiple URLs in parallel (SAFE)

func CheckURLsConcurrently(urls []string) <-chan CheckResult {
	results := make(chan CheckResult, len(urls))
	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(u string) {
			defer wg.Done()
			results <- CheckURL(u)
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
