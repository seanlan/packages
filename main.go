package main

import (
	machinery_config "github.com/RichardKnop/machinery/v2/config"
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/packages/machinery_service"
	"log"
)

func InitCelery() {
	var asyncConfig machinery_config.Config
	var err error
	err = config.GetValue("async_task").Populate(&asyncConfig)
	if err != nil {
		log.Fatal("init celery err: %#v", err)
	}
	machinery_service.Setup(&asyncConfig)
}

func main() {
	config.Setup("")
	logging.Setup(false, "test")
	InitCelery()
	machinery_service.RegisterTasks(map[string]interface{}{
		"add": machinery_service.Add,
	})
	machinery_service.RunWorker("async_service",10)
	//machinery_service.SendTask("add", []tasks.Arg{
	//	tasks.Arg{"", "int64", 10},
	//	tasks.Arg{"", "int64", 10},
	//	tasks.Arg{"", "int64", 10},
	//	tasks.Arg{"", "int64", 10},
	//}, "", 0)
}
