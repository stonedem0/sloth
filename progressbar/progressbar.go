package progressbar

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// PrintProgressBar prints an awesome progress bar
func PrintProgressBar(percent float32) {
	width := 50
	fg := "█"
	bg := "░"
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled
	colors := []uint8{57, 93, 129, 165, 201, 201}
	// fgBar := strings.Repeat(bg, unfilled)
	// println(colors[filled/10], width)
	fgBar := aurora.Index(colors[filled/10], strings.Repeat(bg, unfilled)).String()
	// gBar := aurora.Index(colors[filled/10], strings.Repeat(bg, unfilled)).String()
	bgBar := aurora.Index(colors[filled/10], strings.Repeat(fg, filled)).String()
	// bgBar := strings.Repeat(fg, filled)
	// println(fg)
	// for _, c := range colors {
	// 	println(c)
	// }
	fmt.Printf("\r %s %s %d %s", bgBar, fgBar, aurora.Index(57, int(percent*100)), aurora.Index(57, "%"))
}

// Gradient colors
// 	colors := []uint8{57, 93, 129, 165, 201}
