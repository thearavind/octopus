package main

import (
	"time"
	"github.com/greenac/octopus/logger"
	//"github.com/greenac/octopus/workers"
)

func main() {
	logger.Log("Time:", time.Now().Format(time.UnixDate))
	counter := 1
	otherCounter := 0
	for {
		logger.Log("counter:", counter, "other counter:", otherCounter)
		counter += 1
		otherCounter = counter * counter
		time.Sleep(1 * time.Second)
	}
}
