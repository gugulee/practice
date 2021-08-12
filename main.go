package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler sets up a signal handler that calls the given CancelFunc
// on SIGTERM/SIGINT. If a second signal is caught, the program is terminated
// immediately with exit code 1.
func SetupSignalHandler() {
	close(onlyOneSignalHandler) // panics when called twice
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		fmt.Println("do something")
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()
}

func main() {
	// Set up signals so we handle the first shutdown signal gracefully.
	SetupSignalHandler()
	// SetupSignalHandler()
	for {
		fmt.Println("running...")
		time.Sleep(60 * time.Second)
	}
}
