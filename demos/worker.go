package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		cb()
	}()
}

type Worker struct {
	cancel context.CancelFunc
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) Run() {

	wg := WaitGroupWrapper{}
	ctx, cancel := context.WithCancel(context.Background())
	w.cancel = cancel

	wg.Wrap(func() {
		w.Process1(ctx)
	})

	wg.Wrap(func() {
		w.Process2(ctx)
	})

	wg.Wait()
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
	w.cancel()
}

func (w *Worker) Process1(ctx context.Context) {
	fmt.Println("Process1")
	defer fmt.Println("Process1 end")

	for {
		if w.isDone(ctx) {
			return
		}

		time.Sleep(time.Second / 3)
		fmt.Println("aaaaa")
	}
}

func (w *Worker) Process2(ctx context.Context) {
	fmt.Println("Process2")
	defer fmt.Println("Process2 end")

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
		fmt.Println("w.Run()")
	}()
	for {
		time.Sleep(time.Second * 2)
		w.Stop()
		break
	}
	time.Sleep(time.Second * 3)

	fmt.Println("end....")
}
