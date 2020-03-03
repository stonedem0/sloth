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
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()

	results := make(chan Result)
	go Sloth(urls, *count, results)
	for a := 0; a < len(urls)**count; a++ {
		r := <-results

		if r.Error == nil {
			fmt.Printf("Response %d from %s took %s \n", aurora.Yellow(a), aurora.Green(r.URL), aurora.Magenta(r.Duration))
			continue
		}
		fmt.Printf("Response %d from %s has an error: %s \n", a, r.URL, r.Error)
	}
	close(results)
}

type Result struct {
	Error    error
	Index    int
	Duration time.Duration
	URL      string
}

// Sloth ...
func Sloth(urls []string, count int, res chan Result) {
	var wg sync.WaitGroup
	wg.Add(len(urls) * count)
	for _, val := range urls {
		for a := 0; a < count; a++ {
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
	}

	wg.Wait()
}
