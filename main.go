package main

import (
	"fmt"
	"math/rand"
	"os"

	//. "os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func logGenerator() {
	rand.Seed(time.Now().UnixNano())
	msgs := []string{"An error occured", "Bruteforce attack detected", "Unhandled exception", "Cannot cast type", "Multiple failed login attempts detected"}
	r := rand.Intn(len(msgs))
	log.SetFormatter(&log.JSONFormatter{})
	log.Info(msgs[r])
}

func getNanoSec(startTime time.Time) int64 {
	current := time.Now()
	diff := current.Sub(startTime)
	return diff.Nanoseconds()
}

func initializeLogging(logFile string) {

	/*	var file, err = OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
		if err != nil {
			fmt.Println("Could Not Open Log File : " + err.Error())
		}
		log.SetOutput(file)
	*/
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	initializeLogging("test.log")
	count := 0
	eps, _ := strconv.Atoi(os.Getenv("EPS"))

	startTime := time.Now()
	fmt.Println("started")
	fmt.Println(time.Now())
	for {
		count = count + 1
		if getNanoSec(startTime) >= 1000000000 {
			startTime = time.Now()
			fmt.Println(time.Now())

		}
		if count == eps {
			sec := (getNanoSec(startTime)) * 5
			x := time.Duration(sec) * time.Nanosecond
			time.Sleep(x)
			count = 0
		}
		logGenerator()
	}
}
