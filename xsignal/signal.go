package xsignal

import (
	"os"
	"os/signal"
	"syscall"
)

func Wait() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
