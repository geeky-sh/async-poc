package main

import (
	task "async-poc"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

type TaskRes struct {
	ID       string
	Queue    string
	Type     string
	MaxRetry int
	Retried  int
	Timeout  time.Duration
}

func main() {
	r := chi.NewRouter()

	c := asynq.NewClient(asynq.RedisClientOpt{Addr: "127.0.0.1:6379"})
	defer c.Close()

	am := asynqmon.New(asynqmon.Options{RootPath: "/monitoring", RedisConnOpt: asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		en := json.NewEncoder(w)
		en.Encode(map[string]string{"msg": "success"})
	})

	r.Get("/enqueue", func(w http.ResponseWriter, r *http.Request) {
		res := []TaskRes{}

		var ti *asynq.TaskInfo
		var t1 *asynq.Task
		var err error

		for i := 1; i < 10; i++ {
			t1, err = task.NewTk1Task(i, time.Now())
			if err != nil {
				log.Fatalf("Could not create task t1: %v", err)
			}

			ti, err = c.Enqueue(t1)
			if err != nil {
				log.Fatalf("Could not enqueue task t1: %v", err)
			}

			res = append(res,
				TaskRes{ID: ti.ID, Queue: ti.Queue, Type: ti.Type, MaxRetry: ti.MaxRetry,
					Retried: ti.Retried, Timeout: ti.Timeout})
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		en := json.NewEncoder(w)
		en.Encode(res)
	})

	r.Get(am.RootPath(), am.ServeHTTP)

	http.ListenAndServe(":3001", r)
}
