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
	columns    [][]string
	rows       string
	length     int
	row        int
	colSizes   []int
	data       [][]string
	bg         uint8
	text       uint8
	padding    int
	headers    []string
	header     string
	headerBg   uint8
	headerText uint8
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
func (t *Table) addRow(columns []string, textColor uint8, bgColor uint8, padding int) {
	t.data = append(t.data, columns)
	t.bg = bgColor
	t.text = textColor
	t.padding = padding
	t.colSizes = make([]int, len(columns))
	t.saveColSizes(t.data)

}

//Adding header
func (t *Table) addHeader(headers []string, textColor uint8, bgColor uint8, padding int) {
	t.headers = append(headers)
	t.headerBg = bgColor
	t.headerText = textColor
	// p := " "
	// rows := ""
	// rows += aurora.Index(textColor, strings.Repeat(pad, padding)).BgIndex(bgColor).String()
	// for i, co := range headers {
	// 	p = aurora.Index(textColor, strings.Repeat(pad, (t.colSizes[i]+padding)-len(co))).BgIndex(bgColor).String()
	// 	rows += aurora.Index(textColor, co).BgIndex(bgColor).Bold().String() + p

	// }
	// rows += "\n"
	// t.rows += rows

}

func (t *Table) printHeader(headers []string, textColor uint8, bgColor uint8, padding int) {
	header := ""
	p := " "
	header += aurora.Index(t.text, strings.Repeat(pad, t.padding)).BgIndex(t.headerBg).String()
	for i, co := range headers {
		p = aurora.Index(t.headerText, strings.Repeat(pad, (t.colSizes[i]+padding)-len(co))).BgIndex(t.headerBg).String()
		header += aurora.Index(t.headerText, co).BgIndex(t.headerBg).Bold().String() + p

	}
	t.header += header
}

func (t *Table) addStars() {
	t.rows += aurora.Index(purple, strings.Repeat("⭑ ", 32)).BgIndex(softPink).String()
	t.rows += "\n"
}

//Printing all rows
func (t *Table) printTable() {
	p := " "
	rows := ""
	t.printHeader(t.headers, t.headerBg, t.headerText, t.padding)

	for _, c := range t.data {
		rows += aurora.Index(t.text, strings.Repeat(pad, t.padding)).BgIndex(t.bg).String()
		for i, co := range c {
			p = aurora.Index(t.text, strings.Repeat(pad, (t.colSizes[i]+t.padding)-len(co))).BgIndex(t.bg).String()
			rows += aurora.Index(t.text, co).BgIndex(t.bg).String() + p
		}
		rows += "\n"
	}
	t.rows += rows
	fmt.Printf("%+2v\n", t.header)
	fmt.Printf("%+2v\n", t.rows)
}

// PrintTable accepts map of results and print a table with it.
func PrintTable(m map[string][]time.Duration, c int) {
	t := Table{}
	data := [][]string{}
	headers := []string{"URL", "average", "min", "max"}
	// test := []string{"meh", "bleh", "dude", "123456789019293"}
	t.addHeader(headers, softPink, purple, 3)

	for k, v := range m {
		runes := []rune(k)
		prettyURL := string(runes[8:])
		t.addRow(calculateResults(v, c, prettyURL), purple, softPink, 3)
		data = append(data, calculateResults(v, c, prettyURL))

	}

	// t.addRow(test, purple, softPink, 3)
	t.printTable()

}

// bg := aurora.Index(purple, strings.Repeat("⭑ ", tableSize/2)).BgIndex(softPink)
// sloth := aurora.Index(softPink, "(ﾉ◕ヮ◕)ﾉ*:・ﾟ✧look at this*:・ﾟ✧")
// results := strings.Repeat(pad, 2) + prettyUrl + strings.Repeat(pad, urlColumnSize-len(k)) + av.String() + min.String() + max.String()
