package display

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// ClearScreen clears the terminal screen
func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// DisplayProgressBar displays a progress bar with the given progress
func DisplayProgressBar(current, total int) {
	// Calculate progress percentage
	progress := float64(current) / float64(total) * 100

	// Create a simple progress bar
	const width = 50
	completed := int(width * progress / 100)
	bar := make([]byte, width)
	for i := 0; i < completed; i++ {
		bar[i] = '='
	}
	if completed < width {
		bar[completed] = '>'
	}
	for i := completed + 1; i < width; i++ {
		bar[i] = ' '
	}

	// Show minimal progress information
	fmt.Printf("\r[%s] %.1f%% Game %d/%d", string(bar), progress, current+1, total)
}

// DisplayStatistics displays training statistics
func DisplayStatistics(gameNum, totalGames int, xWins, oWins, draws int, epsilon float64) {
	// Calculate percentages
	xWinRate := float64(xWins) / float64(gameNum+1) * 100
	oWinRate := float64(oWins) / float64(gameNum+1) * 100
	drawRate := float64(draws) / float64(gameNum+1) * 100

	// Clear screen
	ClearScreen()

	// Display statistics
	fmt.Printf("\n=== Training Statistics (Game %d) ===\n", gameNum+1)
	fmt.Printf("Win Rates:\n")
	fmt.Printf("  X: %.1f%% (%d wins)\n", xWinRate, xWins)
	fmt.Printf("  O: %.1f%% (%d wins)\n", oWinRate, oWins)
	fmt.Printf("  Draws: %.1f%% (%d draws)\n", drawRate, draws)
	fmt.Printf("\nTraining Parameters:\n")
	fmt.Printf("  Epsilon: %.3f\n", epsilon)
	fmt.Println("==========================================\n")

	// Wait for 5 seconds
	time.Sleep(5 * time.Second)
}
