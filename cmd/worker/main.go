package main

import (
	task "async-poc"
	"log"

	"github.com/hibiken/asynq"
)

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "127.0.0.1:6379"},
		asynq.Config{Concurrency: 1000})

	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeTk1, task.HandleTk1Task)
	mux.HandleFunc(task.TypeTk2, task.HandleTk2Task)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("Unable to run server %+v", err)
	}
}
