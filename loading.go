package util

import (
	"fmt"
	"os"
	"os/exec"

	"gorm.io/gorm"
)

// loading bar when sql doing some job
func LoadingBar(execQuery *gorm.DB) {
	totalCount, tempCount := int64(0), int64(0)

	execQuery.Count(&totalCount)

	for {
		// task proccess percentage
		percentage := float64(totalCount-tempCount) / float64(totalCount) * 100

		// print (present num / total num) and percentage
		fmt.Println(totalCount-tempCount, "/", totalCount, percentage, "%")

		// print the loading bar
		for i := 0.0; i < 100; i++ {
			if i < percentage {
				print("*")
				continue
			}

			print("-")
		}

		execQuery.Count(&tempCount)

		// refresh the cmd
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
