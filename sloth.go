package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/logrusorgru/aurora"
)

func main() {
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()
	results := make(chan Result)
	total := len(urls) * *count
	m := map[string][]time.Duration{}
	go Sloth(urls, *count, results)
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)
		if r.Error == nil {
			printProgressBar(float32(a) / float32(total))
			continue
		}

		fmt.Printf("Response %d from %s has an error: %s\n", r.Index, r.URL, r.Error)
	}
	close(results)
	printTable(m, *count)
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

var progressBar string

func init() {
	colors := []uint8{57, 93, 129, 165, 201}
	for _, c := range colors {
		progressBar += aurora.Index(c, "■ ■ ■ ■ ■ ").String()
	}
}
func printProgressBar(percent float32) {
	size := utf8.RuneCountInString(progressBar)
	index := int(float32(size) * float32(percent))
	bar := string([]rune(progressBar)[0:index])
	fmt.Printf("\r %s %%%d", bar, int(percent*100))
}

func printTable(m map[string][]time.Duration, c int) {
	msg := "average respond time (ms)"
	fmt.Printf("\n")
	fmt.Printf("%s", strings.Repeat("╌╌", 28))
	fmt.Printf("\n")
	fmt.Printf("%13s %12s %s %4s", "URL", "┆", msg, "┆")
	fmt.Printf("\n")
	fmt.Printf("%s", strings.Repeat("╌╌", 28))
	fmt.Printf("\n")
	pad := " "
	for k, v := range m {
		var sum time.Duration
		var d time.Duration
		for _, s := range v {
			sum = sum + s
		}
		average := sum / time.Duration(c)
		d, _ = time.ParseDuration(average.String())
		fmt.Printf("  %s %s %d %s %s\n", aurora.Index(46, k), strings.Repeat(pad, 34-len(k)), aurora.Index(198, int(d)/1000000), strings.Repeat(pad, 13), aurora.Index(201, "┆"))
	}
}
