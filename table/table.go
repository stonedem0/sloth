package table

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	softPink = 213
	purple   = 57
)

func calculateResults(sl []time.Duration, c int) (time.Duration, time.Duration, time.Duration) {
	var sum time.Duration
	min := sl[0]
	max := sl[1]
	for _, t := range sl {
		sum = sum + t
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	average := sum / time.Duration(c)
	// min = min / 1000000
	// max = max / 1000000
	return average, min, max
}

// Table is a struct with method
// Method addRow

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {

	// Column sizes
	tableSize := 60
	urlColumnSize := 32
	durationColumnSize := 15
	pad := " "
	bg := aurora.Index(purple, strings.Repeat("⭑ ", tableSize/2)).BgIndex(softPink)
	// sloth := aurora.Index(softPink, "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧look at this*:・ﾟ✧")
	header := strings.Repeat(pad, 4) + "URL" + strings.Repeat(pad, 19) + "average" + strings.Repeat(pad, 8) + "min" + strings.Repeat(pad, 8) + "max" + strings.Repeat(pad, 4)
	// fmt.Printf(" %s\n", sloth)
	fmt.Printf("\n")
	fmt.Printf(" %s\n", aurora.Index(softPink, header).BgIndex(purple).Bold())
	fmt.Printf(" %s\n", bg)
	for k, v := range m {
		runes := []rune(k)
		prettyUrl := string(runes[8:])
		var a time.Duration

		var mi time.Duration
		var ma time.Duration
		av, min, max := calculateResults(v, c)

		a, _ = time.ParseDuration(av.String())
		mi, _ = time.ParseDuration(min.String())
		ma, _ = time.ParseDuration(max.String())
		avP := durationColumnSize - len(strconv.Itoa(int(a)))
		minP := durationColumnSize - len(strconv.Itoa(int(mi)))
		maxP := durationColumnSize - len(strconv.Itoa(int(ma)))

		results := strings.Repeat(pad, 2) + prettyUrl + strings.Repeat(pad, urlColumnSize-len(k)) + av.String() + strings.Repeat(pad, avP) + min.String() + strings.Repeat(pad, minP) + max.String() + strings.Repeat(pad, maxP)
		fmt.Printf(" %s\n", aurora.Index(purple, results).Bold().BgIndex(softPink))

	}
	fmt.Printf(" %s\n", bg)
}
