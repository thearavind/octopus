package main

import (
	"time"
	"fmt"
	"github.com/greenac/octopus/logger"
)

func main() {
	fmt.Println(time.Now().Format(time.UnixDate))
	logger.Log("whatever i want")
	counter := 1
	otherCounter := 0
	for {
		logger.Log("counter:", counter, "other counter:", otherCounter)
		counter += 1
		otherCounter = counter * counter
		time.Sleep(1 * time.Second)
	}
}
