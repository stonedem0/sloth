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
	tmp := make([]time.Duration, len(urls)**count)
	go Sloth(urls, *count, results)
	for a := 0; a < len(urls)**count; a++ {
		r := <-results
		tmp = append(tmp, r.Duration)
		if r.Error == nil {
			fmt.Printf("Response %-5d from %s took %s\n", aurora.Yellow(r.Index), aurora.Green(r.URL), aurora.Magenta(r.Duration))
			continue
		}
		fmt.Printf("Response %d from %s has an error: %s\n", r.Index, r.URL, r.Error)
	}
	close(results)
	var sum time.Duration
	for _, t := range tmp {
		sum = sum + t
	}
	fmt.Printf("Average respond speed: %s\n", aurora.Magenta(sum/time.Duration(*count)))
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
			go func(val string, index int) {
				defer wg.Done()
				start := time.Now()
				r, err := http.Get(val)
				if err != nil {
					res <- Result{Index: index, URL: val, Error: err}
					return
				}
				defer r.Body.Close()
				elapsed := time.Since(start).Round(time.Millisecond)
				res <- Result{Index: index, Duration: elapsed, URL: val}
			}(val, a)
		}
	}
	wg.Wait()
}
