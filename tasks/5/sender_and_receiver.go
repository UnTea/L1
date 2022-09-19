package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// Read is a function that read data from the channel while channel is active
func Read(channel chan string) {
	counter := 1

	for data := range channel {
		fmt.Println("read data: ", data, " ", counter, " time(s)")
		counter++
	}

	fmt.Println("channel closed")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Write 1 arg: time of work in ms")
		return
	}

	execTime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("You need to write a number")
		return
	}

	channel := make(chan string)
	go Read(channel)

	duration := time.Duration(execTime) * time.Millisecond
	timer := time.NewTimer(duration)

	for {
		select {
		case <-timer.C:
			close(channel)

			fmt.Println("time's out")
			time.Sleep(time.Millisecond * 100)

			return
		default:
			channel <- "Hello World!"
		}

		time.Sleep(time.Millisecond * 100)
	}
}
