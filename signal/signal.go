package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandlerFunc func(sig os.Signal) (err error)

var ErrStop = errors.New("stop serve signals")

var handlers = make(map[os.Signal]SignalHandlerFunc)

func init() {
	handlers[syscall.SIGTERM] = sigtermDefaultHandler
}

func sigtermDefaultHandler(sig os.Signal) error {
	fmt.Printf("signal %s received, exiting, bye!", sig)
	return ErrStop
}

func SetSigHandler(handler SignalHandlerFunc, signals ...os.Signal) {
	for _, sig := range signals {
		handlers[sig] = handler
	}
}

// ServeSignals calls handlers for system signals.
func ServeSignals() (err error) {
	signals := make([]os.Signal, 0, len(handlers))
	for sig := range handlers {
		signals = append(signals, sig)
	}

	ch := make(chan os.Signal, 8)
	signal.Notify(ch, signals...)

	for sig := range ch {
		err = handlers[sig](sig)
		if err != nil {
			break
		}
	}

	signal.Stop(ch)

	if err == ErrStop {
		err = nil
	}

	return
}

func main() {
	ch := make(chan struct{}, 1)
	SetSigHandler(func(sig os.Signal) (err error) {
		fmt.Printf("signal %s received\n", sig)
		if sig == syscall.SIGUSR1 {
			ch <- struct{}{}
		}
		return nil
	}, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	go ServeSignals()

	select {
	case <-ch:
		fmt.Println("got signal SIGUSR1")
	}

}
