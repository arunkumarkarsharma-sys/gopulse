package checker

import (
	"net/http"

	"time"
)

type Result struct {
	Url          string
	Status       string
	Responsetime time.Duration
	Error        error
}

func CheckURLsConcurrently(urls []string) []Result {

	resultsChan := make(chan Result)

	for _, Url := range urls {
		go checkURL(Url, resultsChan)
	}
	var results []Result

	for i := 0; i < len(urls); i++ {
		result := <-resultsChan
		results = append(results, result)
	}

	return results
}

func checkURL(url string, ch chan Result) {

	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- Result{
			Url:   url,
			Error: err,
		}
		return
	}

	ch <- Result{
		Url:          url,
		Status:       resp.Status,
		Responsetime: time.Since(start),
	}

}
