package main

import (
	"fmt"
	"github.com/seanlan/packages/config"
)

func main() {
	config.Setup("")
	fmt.Println(config.GlobalConfig)
}
