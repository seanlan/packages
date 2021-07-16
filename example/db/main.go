package main

import (
	"fmt"
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/db"
)

func main() {
	config.Setup("")
	db.Setup(config.GetString("mysql"))
	var cnt int64
	db.DB.Table("tb_good_spu").Count(&cnt)
	fmt.Println("cnt:",cnt)
}
