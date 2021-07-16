package main

import (
	machinery_config "github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/tasks"
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
		log.Fatalf("init celery err: %#v", err)
	}
	machinery_service.Setup(&asyncConfig)
}

func StarWorker()  {
	machinery_service.RegisterTasks(map[string]interface{}{
		"add": machinery_service.Add,
	})
	machinery_service.RunWorker("async_service",10)
}

func TestSender()  {
	machinery_service.SendTask("add", []tasks.Arg{
		tasks.Arg{"", "int64", 10},
		tasks.Arg{"", "int64", 10},
		tasks.Arg{"", "int64", 10},
		tasks.Arg{"", "int64", 10},
	}, "", 0)
}

func main() {
	config.Setup("example/conf.d/conf.yaml")
	logging.Setup(false, "test")
	InitCelery()
	StarWorker()
	//TestSender()
}
