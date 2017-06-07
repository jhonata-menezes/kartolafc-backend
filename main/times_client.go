package main

import (
	"github.com/benmanns/goworker"
	"os"
	"strconv"
	"log"
)

func init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://" + os.Args[1] + ":6379/",
		Connections:    100,
		Queues:         []string{"times"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "kartolafc:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
}

func main() {
	ini, _ := strconv.Atoi(os.Args[2])
	fim, _ := strconv.Atoi(os.Args[3])
	for i:=fim; i>=ini; i-- {
		log.Println("enviando", i)
		goworker.Enqueue(&goworker.Job{
			Queue: "times",
			Payload: goworker.Payload{
				Class: "Id",
				Args: []interface{}{i},
			},
		})
	}

}