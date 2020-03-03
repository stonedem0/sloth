package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"
)

func main() {
	flag.Parse()
	urls := flag.Args()
	results := make(chan Result)
	go Sloth(urls, results)
	for a := 0; a < len(urls); a++ {
		r := <-results
		fmt.Println(aurora.Yellow("Response from: "), aurora.Green(r.URL), aurora.Yellow("took: "), aurora.Magenta(r.Duration))
	}
	close(results)
}

type Result struct {
	Error    error
	Duration time.Duration
	URL      string
}

// Sloth ...
func Sloth(urls []string, res chan Result) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, val := range urls {
		go func(val string) {
			defer wg.Done()
			start := time.Now()

			r, err := http.Get(val)
			if err != nil {
				res <- Result{URL: val, Error: err}
				return
			}
			defer r.Body.Close()
			elapsed := time.Since(start).Round(time.Millisecond)
			res <- Result{Duration: elapsed, URL: val}
		}(val)
	}

	wg.Wait()
}
