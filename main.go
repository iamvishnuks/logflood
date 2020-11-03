package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

func logGenerator() {
	for {
		log.SetFormatter(&log.JSONFormatter{})
		log.Info("Something happened")
		time.Sleep(1 * time.Second)
	}
}
func main() {
	logGenerator()
}
