package main

import (
	"context"
	"fmt"
	"github.com/liyue201/golib/cron"
	"time"
)

func process1(ctx context.Context) {
	fmt.Println("process1 start")
	for i := 0; i < 4; i++ {
		fmt.Println("process1", i)
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("process1 interuption")
			return
		default:
		}
	}
	fmt.Println("process1 end")
}

func process2(ctx context.Context) {
	fmt.Println("process2 start")
	for i := 0; i < 5; i++ {
		fmt.Println("process2", i)
		time.Sleep(time.Second)

		select {
		case <-ctx.Done():
			fmt.Println("process2 interuption")
			return
		default:
		}
	}
	fmt.Println("process2 end")
}

func main() {
	c := cron.NewCron(time.Second)
	c.AddTask(&cron.Task{F: process1, D: time.Second * 2})
	c.AddTask(&cron.Task{F: process2, D: time.Second * 3})

	go func() {
		c.Run()
	}()
	go func() {
		time.Sleep(time.Second * 10)
		c.Stop()
	}()

	time.Sleep(time.Minute * 5)
}
