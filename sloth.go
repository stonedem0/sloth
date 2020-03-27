package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/stonedem0/sloth/progressbar"
	"github.com/stonedem0/sloth/table"
	"github.com/stonedem0/sloth/terminal"
	"github.com/stonedem0/sloth/validator"
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
	for _, u := range urls {
		validator.URLValidator(u)
	}
	results := make(chan Result)
	total := len(urls) * *count
	m := map[string][]time.Duration{}

	// Terminal setup
	terminal.CleanTerminalScreen()
	terminal.MoveCursorUpperLeft()
	terminal.HideCursor()

	go Sloth(urls, *count, results)
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)
		if r.Error == nil {
			// Helps to avoid increasing of progress bar, limits it to one line.
			if total > 100 {
				total = 100
			}
			progressbar.PrintProgressBar(float32(a)/float32(total), int(total))
			continue
		}
		fmt.Printf("Response %d from %s has an error: %s\n", r.Index, r.URL, r.Error)
	}
	close(results)
	terminal.EraseProgressBar()
	terminal.MoveCursorUpperLeft()
	table.PrintTable(m, *count)

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
