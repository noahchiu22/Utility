package util

import (
	"errors"
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

				roundPercentage := math.Round(percentage*math.Pow10(b.Precision)) / math.Pow10(b.Precision)
				remaining := -1.0
				if speed > 0 {
					remaining = (max - current) / speed
				}
				fmt.Printf(" | %.*f%% (%.0f/%.0f, %.1fit/s) - %.1fs", b.Precision, roundPercentage, current, max, speed, remaining)
				fmt.Println()

				if current >= max {
					b.Stop()
					continue
				}

				time.Sleep(time.Duration(1/b.FPS) * time.Second)
				clearScreen()
			}
		}
	}()
}

func (b *Bar) Add(load float64) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.StopChan == nil {
		return errors.New("bar has not been initialized")
	}

	now := time.Now()
	timeDiff := now.Sub(b.lastUpdateTime).Seconds()

	b.Current += load

	if timeDiff > 0 {
		b.Speed = load / timeDiff
	}

	b.lastUpdateTime = now

	return nil
}

func (b *Bar) Refresh(current float64) error {
	b.mutex.Lock()

	defer b.mutex.Unlock()

	if b.StopChan == nil {
		return errors.New("bar has not been initialized")
	}

	now := time.Now()
	timeDiff := now.Sub(b.lastUpdateTime).Seconds()
	if timeDiff > 0 {
		b.Speed = (current - b.Current) / timeDiff
	}

	b.Current = current

	b.lastUpdateTime = now

	return nil
}

func (b *Bar) Reset() error {
	b.mutex.Lock()

	defer b.mutex.Unlock()

	if b.StopChan == nil {
		return errors.New("bar has not been initialized")
	}

	b.Current = 0
	b.Speed = 0

	return nil
}

func (b *Bar) AutoRefresh(fn func() float64, delay time.Duration) {
	// 不需要在這裡上鎖，因為我們要啟動一個獨立的 goroutine
	go func() {
		for {
			select {
			case <-b.StopChan:
				return
			default:
				current := fn()
				b.Refresh(current)
			}

			time.Sleep(delay)
		}
	}()
	<-b.StopChan
}

func (b *Bar) Brange() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	return b.Current < b.Max
}

func (b *Bar) Stop() error {
	b.mutex.Lock()

	defer b.mutex.Unlock()

	if b.StopChan == nil {
		return errors.New("bar has not been initialized")
	}

	if b.stopped {
		return nil
	}

	b.stopped = true
	close(b.StopChan)

	return nil
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
