package checker

import (
	"net/http"
	"time"
)

// CheckResult holds the result of a URL check
type CheckResult struct {
	URL          string
	Status       string
	ResponseTime time.Duration
	Error        error
}

// CheckURL checks a single URL
func CheckURL(url string) CheckResult {
	start := time.Now()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

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

// CheckURLsConcurrently checks multiple URLs in parallel
func CheckURLsConcurrently(urls []string) <-chan CheckResult {
	results := make(chan CheckResult)

	go func() {
		for _, url := range urls {
			go func(u string) {
				results <- CheckURL(u)
			}(url)
		}
	}()

	return results
}
