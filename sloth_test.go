package main

import (
	"bytes"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	var buf bytes.Buffer
	urls := []string{"htt://stonedemo.wtf", "https://apex.sh"}
	Sloth(&buf, urls)
	for _, u := range urls {
		if !strings.Contains(buf.String(), u) {
			t.Fatalf("expected %q in output", u)
		}
	}
}
