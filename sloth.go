package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/stonedem0/sloth/progressbar"
)

// ./pkg/sloth/sloth.go   library
// ./cmd/sloth/main.go    command

// ./sloth.go             library
// ./cmd/sloth/main.go    command   imports github.com/stonedem0/sloth

func main() {
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()
	results := make(chan Result)
	total := len(urls) * *count
	m := map[string][]time.Duration{}
	setupTerminal()
	go Sloth(urls, *count, results)
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)
		if r.Error == nil {
			progressbar.PrintProgressBar(float32(a)/float32(total), int(total))
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

// func init() {
// 	colors := []uint8{57, 93, 129, 165, 201}
// 	for _, c := range colors {
// 		fmt.Println(c)
// 		progressBar += aurora.Index(c, "■■■■■").String()
// 		background += aurora.Index(c, "▢▢▢▢▢").String()
// 	}
// 	fmt.Println(len(progressBar))
// 	fmt.Println(len(background))
// }

func setupTerminal() {
	fmt.Printf("\033[2J")
	fmt.Printf("\033[f")
	fmt.Printf("\n")
	fmt.Printf("\033[?25l")
}

func printTable(m map[string][]time.Duration, c int) {

	// Erase progress bar
	fmt.Printf("\033[1K")
	fmt.Printf("\033[f")

	// Column sizes
	urlColumnSize := 30
	durationColumnSize := 10

	fmt.Printf("\n")
	pad := " "
	header := strings.Repeat(pad, 2) + "average(ms)" + strings.Repeat(pad, 17) + "URL" + strings.Repeat(pad, 7)
	fmt.Printf(" %s\n", aurora.Index(213, header).BgIndex(93).Italic())
	for k, v := range m {
		var sum time.Duration
		var d time.Duration
		for _, s := range v {
			sum = sum + s
		}
		average := sum / time.Duration(c)
		d, _ = time.ParseDuration(average.String())
		p := durationColumnSize - len(strconv.Itoa(int(d)/1000000))
		results := k + strings.Repeat(pad, urlColumnSize-len(k)) + strconv.Itoa(int(d)/1000000) + strings.Repeat(pad, p)
		fmt.Printf(" %s\n", aurora.Index(255, results).BgIndex(57))
	}
}
