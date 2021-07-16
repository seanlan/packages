package main

import (
	"fmt"
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/db"
)

func main() {
	config.Setup("example/conf.d/conf.yaml")
	db.Setup(config.GetString("mysql"))
	var cnt int64
	db.DB.Table("tb_good_spu").Count(&cnt)
	fmt.Println("cnt:",cnt)
}
