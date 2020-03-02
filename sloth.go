package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/logrusorgru/aurora"
)

func main() {
	flag.Parse()
	args := flag.Args()
	results := make(chan Result)
	go Sloth(args, results)
	printResult(args, results)
	// dnsLookup()
}

type Result struct {
	Error    error
	Duration time.Duration
	URL      string
	Header   http.Header
}

// Sloth ...
func Sloth(urls []string, res chan Result) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	//for _, val := range urls {
	val := urls[0]
	go func(val string) {
		defer wg.Done()
		start := time.Now()

		r, err := http.Get(val)
		if err != nil {
			res <- Result{URL: val, Error: err}
			return
		}
		defer r.Body.Close()
		header := r.Header
		elapsed := time.Since(start).Round(time.Millisecond)
		res <- Result{Duration: elapsed, URL: val, Header: header}
	}(val)
	//}
	wg.Wait()
}

// func dnsLookup() {
// 	addr, err := net.LookupIP("https://stonedemo.wtf")
// 	if err != nil {
// 		fmt.Println("Unknown host")
// 	} else {
// 		fmt.Println("IP address: ", addr)
// 	}
// }
func printResult(args []string, res chan Result) {
	for a := 0; a < len(args); a++ {
		r := <-res
		if len(args) > 1 {
			switch args[1] {
			case "-s":
				fmt.Println(aurora.Yellow("Response from: "), aurora.Green(r.URL), aurora.Yellow("took: "), aurora.Magenta(r.Duration))
				return
			case "-h":
				for k, v := range r.Header {
					fmt.Print(aurora.Green(k))
					fmt.Print(" : ")
					fmt.Println(aurora.Magenta(v))
				}
				return
			case "-dns":
				ips, _ := net.LookupIP(r.URL)
				fmt.Println(aurora.Magenta(ips))
				return
			}
		}

	}
	close(res)
}
