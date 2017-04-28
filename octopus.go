package main

import (
	"github.com/greenac/octopus/config"
	"fmt"
)

func main() {
	c := config.Configuration()
	fmt.Println(c)
}
