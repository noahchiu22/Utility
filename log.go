package util

import (
	"fmt"
	"log"
	"os"
	"time"
)

// log by datetime and save in the same directory
func Log(information string, info ...interface{}) {
	t := time.Now().Local().Format("2006-01-02")

	fileName := fmt.Sprint("./log/", t, ".log")

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file open error : %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(information)
	for _, item := range info {
		if item != nil {
			log.Printf("%+v\n", item)
		}
	}
}
