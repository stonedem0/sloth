package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	. "github.com/logrusorgru/aurora"
)

func main() {
	flag.Parse()
	urls := flag.Args()
	results := getTimings(urls)
	fmt.Println("Result: ", Green(<-results))
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
			fmt.Println(Yellow("Response from: "), Green(r.URL), Yellow("took: "), Magenta(r.Duration))
		}(val)

	}
	close(res)
	wg.Wait()
	return res
}
