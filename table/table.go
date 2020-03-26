package table

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {

	// Column sizes
	urlColumnSize := 30
	durationColumnSize := 10

	pad := " "
	header := strings.Repeat(pad, 2) + "URL" + strings.Repeat(pad, 17) + "average(ms)" + strings.Repeat(pad, 7)
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
