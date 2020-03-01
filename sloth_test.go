package main

import (
	"testing"
)

func Test(t *testing.T) {
	urls := []string{"http://stonedemo.wtf", "https://apex.sh"}
	results := make(chan Result)
	go Sloth(urls, results)
	for a := 0; a < len(urls); a++ {
		r := <-results
		if r.Error != nil {
			t.Fatalf("error: %s", r.Error)
			continue
		}
	}
	close(results)
}
