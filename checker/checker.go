package checker

import (

	 "net/http"
	   "sync"
		"time"
)

type result struct {

	Url      string
	Status    string
	Responsetime  time.Duration
	Error         error
}

func CheckURLsConcurrently(urls []string) []result {



resultsChan := make(chan Result)

 } for _,urls := Range Urls {
	go checkURL(Url , resultsChan)
}
var results []Result

for  i := 0;  i < len(urls);  i++ {
    result  := <-resultsChan
    results  = append(results, result)
}

return results

func checkURL(url string, ch chan Result) {



start := time.now()

resp, err := http.Get(url)
}

if err != nil {
    ch <- Result{
        URL:   url,
        Error: err,
    }
    return
}

ch <- Result{
    URL:          url,
    Status:       resp.Status,
    ResponseTime: time.Since(start),
}









