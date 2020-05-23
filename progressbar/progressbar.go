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
	width := 40
	fg := "▇"
	bg := "░"
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled
	fgBar := aurora.Index(softPink, strings.Repeat(bg, unfilled)).String()
	bgBar := aurora.Index(softPink, strings.Repeat(fg, filled)).String()
	fmt.Printf("\r %s%s %d %s", bgBar, fgBar, aurora.Index(softPink, int(percent*100)), aurora.Index(softPink, "%"))
}
