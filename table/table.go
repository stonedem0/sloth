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
	sloth := "🦥"
	_ = sloth
	header := strings.Repeat(pad, 4) + "URL" + strings.Repeat(pad, 20) + "average(ms)" + strings.Repeat(pad, 4)
	fmt.Printf(" %s\n", aurora.Index(213, header).BgIndex(93).Bold())
	fmt.Println("", aurora.Index(57, "▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚▞▚"))
	for k, v := range m {
		var sum time.Duration
		var d time.Duration
		for _, s := range v {
			sum = sum + s
		}
		average := sum / time.Duration(c)
		d, _ = time.ParseDuration(average.String())
		p := durationColumnSize - len(strconv.Itoa(int(d)/1000000))
		results := strings.Repeat(pad, 2) + k + strings.Repeat(pad, urlColumnSize-len(k)) + strconv.Itoa(int(d)/1000000) + strings.Repeat(pad, p)
		fmt.Printf(" %s\n", aurora.Index(154, results).BgIndex(57))
	}
}
