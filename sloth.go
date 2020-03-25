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
)

func main() {
	count := flag.Int("count", 10, "number of requests")
	flag.Parse()
	urls := flag.Args()
	results := make(chan Result)
	total := len(urls) * *count
	m := map[string][]time.Duration{}
	go Sloth(urls, *count, results)
	fmt.Printf("\n")
	for a := 0; a < total; a++ {
		r := <-results
		m[r.URL] = append(m[r.URL], r.Duration)
		if r.Error == nil {
			printProgressBar(float32(a)/float32(total), int(total))
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
var background string

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

func printProgressBar(percent float32, width int) {
	fg := "█"
	bg := "░"
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled - 1
	bar := strings.Repeat(bg, unfilled)
	back := strings.Repeat(fg, filled)
	// fmt.Printf("\r %s %s", back, bar)
	fmt.Printf("\r %s %s", aurora.Index(57, back), aurora.Index(57, bar))
}

func printTable(m map[string][]time.Duration, c int) {
	// tableSize := 60
	urlColumnSize := 30
	durationColumnSize := 10
	fmt.Printf("\n")
	fmt.Printf("\n")
	msg := "average(ms)"
	fmt.Printf("%s", strings.Repeat("╌*", 23))
	fmt.Printf("\n")
	fmt.Printf("%10s %30s %4s", "URL", msg, "┆")
	fmt.Printf("\n")
	fmt.Printf("%s %s", strings.Repeat("╌", 44), "╋")
	fmt.Printf("\n")
	for k, v := range m {
		var sum time.Duration
		var d time.Duration
		for _, s := range v {
			sum = sum + s
		}
		pad := " "
		average := sum / time.Duration(c)
		d, _ = time.ParseDuration(average.String())
		p := durationColumnSize - len(strconv.Itoa(int(d)/1000000))
		fmt.Printf(" %s %s %d %s %s\n", aurora.Index(46, k), strings.Repeat(pad, urlColumnSize-len(k)), aurora.Index(198, int(d)/1000000), strings.Repeat(pad, p), aurora.Index(201, "┆"))
	}
}
