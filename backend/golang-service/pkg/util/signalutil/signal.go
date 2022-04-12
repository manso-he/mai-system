package signalutil

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler(ch chan<- struct{}, fn func()) {
	c := make(chan os.Signal, 5)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-c
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(c)
			fn()
			ch <- struct{}{}
			return
		}
	}
}
