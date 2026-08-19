package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func mainsignal() {

	pid := os.Getpid()
	fmt.Println("Proccess id >>", pid)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify channel on interrupt or terminate signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) //syscall.SIGUSR1)

	go func() {
		sig := <-sigs
		fmt.Println("Received signal: ", sig)
		done <- true
	}()

	go func() {

		for {
			select {
			case <-done:
				fmt.Println("Stopping work due to signal.")
				// os.Exit(0)
				return
			default:
				fmt.Println("working...")
				time.Sleep(time.Second)
			}
		}

		// sig := <-sigs
		// switch sig {
		// case syscall.SIGINT:
		// 	fmt.Println("Received signal interrupt")
		// case syscall.SIGTERM:
		// 	fmt.Println("Received signal terminate")
		// case syscall.SIGHUP:
		// 	fmt.Println("Received signal hangup")
		// 	// case syscall.SIGUSR1: //only in linux/mac
		// 	// 	fmt.Println("Received signal USR1 ")
		// 	// 	fmt.Println("user define function is executed")
		// }
		// fmt.Println("exit")
		// os.Exit(0)

	}()

	// fmt.Println("working....")
	for {
		time.Sleep(time.Second)
	}
}

// tasklist - List of all process on windows
// taskki; /F /PID <>PIS
