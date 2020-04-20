package table

import (
	"fmt"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

const (
	softPink = 213
	purple   = 57
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

// Table is a struct which contains all columns and rows for table
type Table struct {
	columns []string
	rows    string
}

//Adding row to a table
func (t *Table) addRow(columns []string, textColor uint8, bgColor uint8) {
	pad := " "
	columLength := 24
	for _, v := range columns {
		v += aurora.Index(textColor, strings.Repeat(pad, columLength-len(v))).BgIndex(bgColor).String()
		t.rows += aurora.Index(textColor, v).BgIndex(bgColor).String()
	}
	t.rows += "\n"
}

//Adding header
func (t *Table) addHeader(headers []string, textColor uint8, bgColor uint8) {
	pad := " "
	columLength := 24
	for _, v := range headers {
		v += aurora.Index(textColor, strings.Repeat(pad, columLength-len(v))).BgIndex(bgColor).String()
		t.rows += aurora.Index(textColor, v).Bold().BgIndex(bgColor).String()
	}
	t.rows += "\n"
}

func (t *Table) addStars() {
	t.rows += aurora.Index(purple, strings.Repeat("⭑ ", 10)).BgIndex(softPink).String()
	t.rows += "\n"
}

//Printing all rows
func (t *Table) printTable() {
	fmt.Printf("%+v\n", t.rows)
}

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {
	t := Table{}
	headers := []string{"URL", "average", "min", "max"}
	t.addRow(headers, softPink, purple)
	// t.addStars()

	for k, v := range m {
		runes := []rune(k)
		prettyURL := string(runes[8:])
		results := calculateResults(v, c, prettyURL)

		t.addRow(results, purple, softPink)
	}

	t.printTable()
}

// bg := aurora.Index(purple, strings.Repeat("⭑ ", tableSize/2)).BgIndex(softPink)
// sloth := aurora.Index(softPink, "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧look at this*:・ﾟ✧")
// results := strings.Repeat(pad, 2) + prettyUrl + strings.Repeat(pad, urlColumnSize-len(k)) + av.String() + min.String() + max.String()
