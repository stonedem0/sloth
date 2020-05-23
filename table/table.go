package table

import (
	"time"

	"github.com/stonedem0/roti"
)

const (
	softPink = 213
	purple   = 57
	pad      = " "
)

func calculateResults(t []time.Duration, c int, url string) []string {
	var sum time.Duration
	results := []string{}
	min := t[0]
	max := t[1]
	for _, v := range t {
		sum = sum + v
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	average := sum / time.Duration(c)
	results = append(results, url, average.String(), min.String(), max.String())
	return results
}

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {
	t := roti.Table{}
	headers := []string{"URL", "average", "min", "max"}
	t.AddHeader(headers, softPink, purple, 3)
	for k, v := range m {
		runes := []rune(k)
		prettyURL := string(runes[8:])
		t.AddRow(calculateResults(v, c, prettyURL), purple, softPink, 3)

	}
	println(t.String())

}
