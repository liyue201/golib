package xsync

import (
	"fmt"
	"runtime/debug"
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		if err := recover(); err != nil {
			fmt.Printf("%v\n", err)
			fmt.Printf("%s\n", string(debug.Stack()))
		}
		cb()
	}()
}
