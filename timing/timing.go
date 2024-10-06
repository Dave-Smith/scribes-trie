package timing

import (
	"fmt"
	"time"
)

func Track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func Duration(msg string, start time.Time) {
	fmt.Printf("%v: %v\n", msg, time.Since(start))
}
