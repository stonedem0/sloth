package progressbar

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

const (
	softPink = 213
	purple   = 57
)

// PrintProgressBar prints an awesome progress bar
func PrintProgressBar(percent float32) {

	width := 50
	fg := "█"
	bg := "░"
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled
	fgBar := aurora.Index(softPink, strings.Repeat(bg, unfilled)).String()
	bgBar := aurora.Index(softPink, strings.Repeat(fg, filled)).String()
	fmt.Printf("\r %s %s %d %s", bgBar, fgBar, aurora.Index(softPink, int(percent*100)), aurora.Index(softPink, "%"))
}

// Gradient colors
// 	colors := []uint8{57, 93, 129, 165, 201m 201}
// fgBar := aurora.Index(colors[filled/10], strings.Repeat(bg, unfilled)).String()
// bgBar := aurora.Index(colors[filled/10], strings.Repeat(fg, filled)).String()
