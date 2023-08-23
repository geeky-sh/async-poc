package task

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeTk1 = "type:tk1"
	TypeTk2 = "type:tk2"
)

type Payload struct {
	Counter int       `json:"counter"`
	SentAt  time.Time `json:"sent_at"`
}

// tasks

func NewTk1Task(counter int, sentAt time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(Payload{Counter: counter, SentAt: sentAt})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeTk1, payload), nil
}

func NewTk2Task(counter int, sentAt time.Time) (*asynq.Task, error) {
	payload, err := json.Marshal(Payload{Counter: counter, SentAt: sentAt})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeTk2, payload), nil
}

// end: tasks

func HandleTk1Task(ctx context.Context, t *asynq.Task) error {
	res := Payload{}
	if err := json.Unmarshal(t.Payload(), &res); err != nil {
		return err
	}

	fmt.Printf("Started Task Tk1 with Payload %+v at time %s\n", res, time.Now().Format("01-02-2006 15:04:05"))
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished Task Tk1 with Payload %+v at time %s\n", res, time.Now().Format("01-02-2006 15:04:05"))

	return nil
}

func HandleTk2Task(ctx context.Context, t *asynq.Task) error {
	res := Payload{}
	if err := json.Unmarshal(t.Payload(), &res); err != nil {
		return err
	}

	fmt.Printf("Started Task Tk2 with Payload %+v\n", res)
	time.Sleep(3 * time.Second)
	fmt.Printf("Finished Task Tk2 with Payload %+v\n", res)

	return nil
}
