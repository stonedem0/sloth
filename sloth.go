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
	getTimings(urls)
}

type Result struct {
	Error    error
	Duration time.Duration
	URL      string
}

func getTimings(urls []string) <-chan Result {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	res := make(chan Result)
	for _, val := range urls {
		go func(val string) {
			defer wg.Done()
			start := time.Now()
			_, err := http.Get(val)
			if err != nil {
				panic(err)
			}
			elapsed := time.Since(start).Round(time.Millisecond)
			r := Result{Duration: elapsed, URL: val}
			fmt.Println(aurora.Yellow("Response from: "), aurora.Green(r.URL), aurora.Yellow("took: "), aurora.Magenta(r.Duration))
		}(val)

	}
	close(res)
	wg.Wait()
	return res
}
