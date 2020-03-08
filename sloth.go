package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"
)

func main() {
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()
	results := make(chan Result)
	total := len(urls) * *count
	// colors := []uint8{57, 93, 129, 165, 201}
	m := map[string][]time.Duration{}
	go Sloth(urls, *count, results)
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)
		if r.Error == nil {
			progressBar(float32(a)/float32(total), 20)
			continue
		}
		fmt.Printf("Response %d from %s has an error: %s\n", r.Index, r.URL, r.Error)
	}
	close(results)
	fmt.Printf("\n")
	msg := "Average respond time for"
	for k, v := range m {
		var sum time.Duration
		for _, s := range v {
			sum = sum + s
		}
		fmt.Printf("%s %s: %s\n", aurora.Yellow(msg), aurora.Green(k), aurora.Magenta(sum/time.Duration(*count)))
	}
}

// Result ...
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

func printProgressBar(percent float32, width int) {
	s := int(float32(width) * float32(percent))
	fmt.Printf("\r  %s %%%d", aurora.Index(200, strings.Repeat("■ ", s)), int(percent*100))
}
