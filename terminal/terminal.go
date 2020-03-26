package terminal

import "fmt"

//CleanTerminalScreen just cleans current terminal screen before progress bar printing.
func CleanTerminalScreen() {
	fmt.Printf("\033[2J")
}

//MoveCursorUpperLeft moves cursor so t result able will be printed on upper left corner
func MoveCursorUpperLeft() {
	fmt.Printf("\033[f")
}

// HideCursor hides cursor during progress bar and result printing
func HideCursor() {
	fmt.Printf("\033[?25l")
}

//ShowCursor brings cursor back
func ShowCursor() {
	fmt.Printf("\033[?25h")
}

//EraseProgressBar erases printed progressBar
func EraseProgressBar() {
	fmt.Printf("\033[1K")
}
