package task_queue

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"time"
)

var Celery *machinery.Server

// 初始化Celery
func Setup(celeryServerConfig config.Config) {
	var err error
	Celery, err = machinery.NewServer(&celeryServerConfig)
	if err != nil {
		panic(err)
	}
}

//注册任务
func RegisterTasks(tasks map[string]interface{}) {
	var err error
	err = Celery.RegisterTasks(tasks)
	if err != nil {
		panic(err)
	}
}

//执行任务监听
func RunWorker(consumerTag string, concurrency int, queueName string) {
	worker := Celery.NewCustomQueueWorker(consumerTag, concurrency,queueName)
	err := worker.Launch()
	if err != nil {
		panic(err)
	}
}

//发送任务
func SendTask(taskName string, args []tasks.Arg,routingKey string, delayed int) (*result.AsyncResult, error) {
	eta := time.Now().Add(time.Second * time.Duration(delayed))
	signature := &tasks.Signature{
		RoutingKey:routingKey,
		Name: taskName,
		ETA:  &eta, //定时
		Args: args,
	}
	return Celery.SendTask(signature)
}
