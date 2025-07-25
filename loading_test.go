package util

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	bar := Bar{
		Max:  100,
		Fill: "😂",
	}
	bar.Init()

	for bar.Brange() {
		bar.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
	// bar.AutoRefresh(func() float64 {
	// 	return bar.Current + 1
	// }, 100*time.Millisecond)
	// <-bar.StopChan
}
