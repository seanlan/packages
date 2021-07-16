package main

import (
	"fmt"
	"github.com/seanlan/packages/config"
)

func main() {
	config.Setup("example/conf.d/conf.yaml")
	fmt.Println(config.GlobalConfig)
}
