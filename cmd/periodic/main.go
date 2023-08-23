package main

import (
	task "async-poc"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

func main() {
	sch := asynq.NewScheduler(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}, nil)

	tk, err := task.NewTk1Task(5, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	e, err := sch.Register("@every 15s", tk)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Registered entry %q", e)

	if err := sch.Run(); err != nil {
		log.Fatal(err)
	}
}
