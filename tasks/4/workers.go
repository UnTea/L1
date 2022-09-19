package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Worker is a storage that contained worker id
type Worker struct {
	id int
}

// print is a function that prints the worker(by id) and his value
func (worker *Worker) print(channelInteger chan int) {
	for {
		fmt.Println("Worker", worker.id, "value from the channel: ", <-channelInteger)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	channelInteger := make(chan int)
	exitChannel := make(chan os.Signal)

	// Quit without error when user presses Ctrl+C
	signal.Notify(exitChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Print("Enter the workers count: ")
	var workerCount int

	_, err := fmt.Scanln(&workerCount)
	if err != nil || workerCount < 1 {
		fmt.Printf("Error occurred while inputing workers count: err=%v, workers=%v", err, workerCount)
	}

	// Creates workers and waiting for input value
	for i := 1; i <= workerCount; i++ {
		worker := Worker{
			id: i,
		}

		go worker.print(channelInteger)
	}

	result := 1

	for {
		select {
		case channelInteger <- result:
			result++
		case <-exitChannel:
			fmt.Println("Program stopped")
			os.Exit(0)
		}
	}
}
