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
	// if width > 100 {
	// 	width = 100
	// }
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled
	fgBar := strings.Repeat(bg, unfilled)
	bgBar := strings.Repeat(fg, filled)
	fmt.Printf("\r %s %s %d %s", aurora.Index(57, bgBar), aurora.Index(57, fgBar), aurora.Index(57, int(percent*100)), aurora.Index(57, "%"))
}

// Gradient colors
// 	colors := []uint8{57, 93, 129, 165, 201}
