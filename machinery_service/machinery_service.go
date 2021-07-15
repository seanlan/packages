package machinery_service

import (
	"github.com/RichardKnop/machinery/v2"
	"github.com/RichardKnop/machinery/v2/backends/result"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/tasks"
	"time"
)

var Server *machinery.Server

// Setup 初始化Celery
func Setup(conf *config.Config) {
	brokerServer, _ := BrokerFactory(conf)
	backendServer, _ := BackendFactory(conf)
	lock, _ := LockFactory(conf)
	Server = machinery.NewServer(conf, brokerServer, backendServer, lock)
}

// RegisterTasks 注册任务
func RegisterTasks(tasks map[string]interface{}) {
	var err error
	err = Server.RegisterTasks(tasks)
	if err != nil {
		panic(err)
	}
}

// RunWorker 执行任务监听
func RunWorker(consumerTag string, concurrency int) {
	conf := Server.GetConfig()
	worker := Server.NewCustomQueueWorker(
		consumerTag,
		concurrency,
		conf.DefaultQueue)
	err := worker.Launch()
	if err != nil {
		panic(err)
	}
}

// SendTask 发送任务
func SendTask(taskName string, args []tasks.Arg, routingKey string, delayed int) (*result.AsyncResult, error) {

	signature := &tasks.Signature{
		RoutingKey: routingKey,
		Name:       taskName,
		Args:       args,
	}
	if delayed > 0 {
		eta := time.Now().Add(time.Second * time.Duration(delayed))
		signature.ETA = &eta
	}
	return Server.SendTask(signature)
}
