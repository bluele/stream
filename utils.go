package stream

import (
	"os"
)

func hasStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}
