package progressbar

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// PrintProgressBar prints an awesome progress bar

func PrintProgressBar(percent float32, width int) {
	fg := "█"
	bg := "░"
	filled := int(float32(width) * float32(percent))
	unfilled := width - filled - 1
	fgBar := strings.Repeat(bg, unfilled)
	bgBar := strings.Repeat(fg, filled)
	fmt.Printf("\r %s %s %d %s", aurora.Index(57, bgBar), aurora.Index(57, fgBar), aurora.Index(57, int(percent*100)), aurora.Index(57, "%"))
}
