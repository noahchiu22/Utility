package util

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

type Bar struct {
	Title          string
	Max            float64
	Current        float64
	Precision      int
	Length         float64
	Speed          float64
	Fill           string
	Empty          string
	FPS            float64
	StopChan       chan bool
	stopped        bool
	mutex          sync.Mutex
	startTime      time.Time
	lastUpdateTime time.Time
}

func (b *Bar) Init() {
	b.stopped = false
	b.StopChan = make(chan bool)
	b.Current = 0
	b.startTime = time.Now()
	b.lastUpdateTime = b.startTime

	// default value
	if b.Title == "" {
		b.Title = "Progress"
	}
	if b.Max == 0 {
		b.Max = 100
	}
	if b.Length == 0 {
		b.Length = 20
	}
	if b.Fill == "" {
		b.Fill = "◼︎"
	}
	if b.Empty == "" {
		b.Empty = " "
	}
	if b.FPS == 0 {
		b.FPS = 15
	}

	go func() {
		for {
			select {
			case <-b.StopChan:
				return
			default:
				b.mutex.Lock()
				// task process percentage
				percentage := (b.Current / b.Max) * 100
				current := b.Current
				max := b.Max
				speed := b.Speed
				b.mutex.Unlock()

				// print title
				fmt.Printf("%s:| ", b.Title)

				// print the loading bar (default 20 characters length, more suitable for display)
				for i := 0.0; i < b.Length; i++ {
					if i < (percentage/100)*b.Length {
						fmt.Printf(b.Fill)
					} else {
						fmt.Printf(b.Empty)
					}
				}

				// print (current num / total num) and percentage
				roundNum := math.Round(percentage*math.Pow10(b.Precision)) / math.Pow10(b.Precision)
				remaining := -1.0
				if speed > 0 {
					remaining = (max - current) / speed
				}
				fmt.Printf(" | %.*f%% (%.0f/%.0f, %.1fit/s) - %.1fs", b.Precision, roundNum, current, max, speed, remaining)
				fmt.Println()

				time.Sleep(time.Duration(1/b.FPS) * time.Second)
				clearScreen()
			}
		}
	}()
}

func (b *Bar) Update(current float64) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	now := time.Now()
	timeDiff := now.Sub(b.lastUpdateTime).Seconds()

	b.Current += current

	if timeDiff > 0 {
		b.Speed = current / timeDiff
	}

	b.lastUpdateTime = now
}

func (b *Bar) Stop() {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.stopped {
		return
	}

	b.stopped = true
	close(b.StopChan)
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
