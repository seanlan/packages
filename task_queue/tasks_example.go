package task_queue

import (
	"log"
)

func Add(args ...int64) (int64, error) {
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	return sum, nil
}

func PrintOK() error {
	log.Println("-------OK-------")
	return nil
}

func SayHi(arg string) error {
	log.Println("Hi, i'm " + arg)
	return nil
}
