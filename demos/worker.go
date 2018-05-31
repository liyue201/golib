package main

import (
	"context"
	"fmt"
	"time"
)

type Worker struct {
	stop chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		stop: make(chan struct{}),
	}
}

func (w *Worker) Run() {

	ctx, cancel := context.WithCancel(context.Background())
	go w.ProcessBlock(ctx)
	go w.ProcessMempoolTransaction(ctx)

	for {
		select {
		case <-w.stop:
			cancel()
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

func (w *Worker) isDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
	return false
}

func (w *Worker) Stop() {
	w.stop <- struct{}{}
}

func (w *Worker) ProcessBlock(ctx context.Context) {
	fmt.Println("ProcessBlock")
	defer fmt.Println("ProcessBlock end")

	for {
		if w.isDone(ctx) {
			return
		}

		time.Sleep(time.Second / 3)
		fmt.Println("aaaaa")
	}
}

func (w *Worker) ProcessMempoolTransaction(ctx context.Context) {
	fmt.Println("ProcessMempoolTransaction")
	defer fmt.Println("ProcessMempoolTransaction end")

	for {
		if w.isDone(ctx) {
			return
		}
		time.Sleep(time.Second / 3)
		fmt.Println("bbbbbb")
	}
}

func main() {
	w := NewWorker()
	go func() {
		w.Run()
	}()
	for {
		time.Sleep(time.Second * 2)
		w.Stop()
		break
	}
	time.Sleep(time.Second * 3)

	fmt.Println("end....")
}
