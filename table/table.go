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

// Table is a struct which contains all columns and rows for table
type Table struct {
	columns  [][]string
	rows     string
	length   int
	row      int
	colSizes []int
	data     [][]string
}

func (t *Table) saveColSizes(columns [][]string) []int {
	for _, c := range columns {
		for k, co := range c {
			if t.colSizes[k] < len(co) {
				t.colSizes[k] = len(co)
			}
		}
	}
	return t.colSizes
}

//Adding row to a table
func (t *Table) addRow(data [][]string, textColor uint8, bgColor uint8, padding int) {
	p := " "
	rows := ""
	for _, c := range data {
		rows += aurora.Index(textColor, strings.Repeat(pad, padding)).BgIndex(bgColor).String()
		for i, co := range c {
			p = aurora.Index(textColor, strings.Repeat(pad, (t.colSizes[i]+padding)-len(co))).BgIndex(bgColor).String()
			rows += aurora.Index(textColor, co).BgIndex(bgColor).String() + p
		}
		rows += "\n"
	}
	t.rows += rows

}

//Adding header
func (t *Table) addHeader(headers []string, textColor uint8, bgColor uint8, padding int) {
	p := " "
	rows := ""
	rows += aurora.Index(textColor, strings.Repeat(pad, padding)).BgIndex(bgColor).String()
	for i, co := range headers {
		p = aurora.Index(textColor, strings.Repeat(pad, (t.colSizes[i]+padding)-len(co))).BgIndex(bgColor).String()
		rows += aurora.Index(textColor, co).BgIndex(bgColor).String() + p

	}
	rows += "\n"
	t.rows += rows

}

func (t *Table) addStars() {
	t.rows += aurora.Index(purple, strings.Repeat("⭑ ", 32)).BgIndex(softPink).String()
	t.rows += "\n"
}

//Printing all rows
func (t *Table) printTable() {
	fmt.Printf("%+2v\n", t.rows)
}

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {
	t := Table{}
	size := 0
	for k, _ := range m {
		if len(k) > 0 {
			size++
		}
	}
	t.colSizes = make([]int, size)
	// results := []string{}
	data := [][]string{}
	headers := []string{"URL", "average", "min", "max"}

	for k, v := range m {
		runes := []rune(k)
		prettyURL := string(runes[8:])
		// results = calculateResults(v, c, prettyURL)
		data = append(data, calculateResults(v, c, prettyURL))
		// _ = results

	}

	t.saveColSizes(data)
	t.addHeader(headers, softPink, purple, 3)
	t.addRow(data, purple, softPink, 3)
	t.printTable()

}

// bg := aurora.Index(purple, strings.Repeat("⭑ ", tableSize/2)).BgIndex(softPink)
// sloth := aurora.Index(softPink, "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧look at this*:・ﾟ✧")
// results := strings.Repeat(pad, 2) + prettyUrl + strings.Repeat(pad, urlColumnSize-len(k)) + av.String() + min.String() + max.String()
