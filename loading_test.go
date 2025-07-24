package util

import (
	"math/rand/v2"
	"testing"
	"time"
)

func TestBar(t *testing.T) {
	max := 100.0
	bar := Bar{
		Max:  100,
		Fill: "😂",
	}

	bar.Init()
	for i := 0.0; i < max; i++ {
		bar.Update(0.8 + rand.Float64()*0.4)
		time.Sleep(100 * time.Millisecond)
	}
	bar.Stop()
}
