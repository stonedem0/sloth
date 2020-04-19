package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/stonedem0/sloth/progressbar"
	"github.com/stonedem0/sloth/table"
	"github.com/stonedem0/sloth/terminal"
	"github.com/stonedem0/sloth/validator"

	"github.com/logrusorgru/aurora"
)

//TODO:
// better error hangling

// ./pkg/sloth/sloth.go   library
// ./cmd/sloth/main.go    command

// ./sloth.go             library
// ./cmd/sloth/main.go    command   imports github.com/stonedem0/sloth

func main() {
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		slothMsg := "Hi. It's Sloth. Looks like you're trying to talk to me. Try some of these commands:\n"
		commads := " • [urls]: list of URls for testing "
		fmt.Printf("%s %s", aurora.Index(57, slothMsg), aurora.Index(201, commads))
		os.Exit(1)
	}
	for _, u := range urls {
		validator.URLValidator(u)
	}
	results := make(chan Result)
	total := len(urls) * *count
	m := map[string][]time.Duration{}
	errors := map[string][]error{}

	// Terminal setup
	terminal.CleanTerminalScreen()
	terminal.MoveCursorUpperLeft()
	terminal.HideCursor()

	go Sloth(urls, *count, results)
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)

		if r.Error == nil {
			progressbar.PrintProgressBar(float32(a) / float32(total))
			continue
		}
		// errors = append(errors, r.Error)
		errors[r.URL] = append(errors[r.URL], r.Error)
		// fmt.Printf("Response %d from %s has an error: %s\n", r.Index, r.URL, r.Error)
	}
	close(results)
	terminal.EraseProgressBar()
	terminal.MoveCursorUpperLeft()
	table.PrintTable(m, *count)
	for url, e := range errors {
		for _, err := range e {
			fmt.Printf("%s has an error:\n %s\n", aurora.Index(118, url), aurora.Index(197, err))
		}
	}
	//Show cursor
	terminal.ShowCursor()
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
	// cl := &http.Client{
	// 	Transport: &http.Transport{MaxConnsPerHost: 10},
	// }
	// tr := &http.Transport{
	// 	MaxIdleConns:        20000,
	// 	MaxIdleConnsPerHost: 20000,
	// }
	// cl := &http.Client{Transport: tr}
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
