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

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {

	// Column sizes
	tableSize := 42
	urlColumnSize := 36
	durationColumnSize := 12
	pad := " "
	bg := aurora.Index(purple, strings.Repeat("⭑ ", tableSize/2)).BgIndex(softPink)
	sloth := aurora.Index(softPink, "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧look at this*:・ﾟ✧")
	header := strings.Repeat(pad, 4) + "URL" + strings.Repeat(pad, 20) + "average(ms)" + strings.Repeat(pad, 4)
	fmt.Printf(" %s\n", sloth)
	fmt.Printf("\n")
	fmt.Printf(" %s\n", aurora.Index(softPink, header).BgIndex(purple).Bold())
	fmt.Printf(" %s\n", bg)
	for k, v := range m {
		var sum time.Duration
		var d time.Duration
		for _, s := range v {
			sum = sum + s
		}
		runes := []rune(k)
		prettyUrl := string(runes[8:])
		average := sum / time.Duration(c)
		d, _ = time.ParseDuration(average.String())
		p := durationColumnSize - len(strconv.Itoa(int(d)/1000000))
		results := strings.Repeat(pad, 2) + prettyUrl + strings.Repeat(pad, urlColumnSize-len(k)) + strconv.Itoa(int(d)/1000000) + strings.Repeat(pad, p)
		fmt.Printf(" %s\n", aurora.Index(purple, results).Bold().BgIndex(softPink))

	}
	fmt.Printf(" %s\n", bg)
}
