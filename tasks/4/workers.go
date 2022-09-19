package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Worker struct {
	id int
}

func (worker *Worker) print(channelInteger chan int) {
	for {
		fmt.Println("Worker", worker.id, "value from the channel: ", <-channelInteger)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	channelInteger := make(chan int)
	exitChannel := make(chan os.Signal)

	signal.Notify(exitChannel, os.Interrupt, syscall.SIGTERM)

	fmt.Print("Enter the workers count: ")
	var workerCount int

	_, err := fmt.Scanln(&workerCount)
	if err != nil || workerCount < 1 {
		fmt.Printf("Error occurred while inputing workers count: err=%v, workers=%v", err, workerCount)
	}

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
