package main

import (
	"context"
	"fmt"
	"time"
)

func StopByCancelContext(channel chan string) {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case channel <- "there should be data here":
				fmt.Println("sending data: ")
			case <-ctx.Done():
				close(channel)

				return
			}

			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		time.Sleep(time.Second * 1)
		cancel()
	}()

	for data := range channel {
		fmt.Printf("\treceive data: %v\n", data)
	}
}

func StopByCloseChannel(channel chan string) {
	go func() {
		for {
			v, ok := <-channel
			if !ok {
				fmt.Println("channel is closed")

				return
			}

			fmt.Println(v)
		}
	}()

	channel <- "some"
	channel <- "data"

	close(channel)
	time.Sleep(3 * time.Second)
}

func StopByPanic(channel chan string) {
	go func() {
		for {
			channel <- "Hello"

			time.Sleep(3 * time.Second)
			panic("Stop by panic")

			channel <- "world"
		}
	}()
}

func main() {
	channel := make(chan string)

	StopByCancelContext(channel)
	//StopByCloseChannel(channel)
	//StopByPanic(channel)
}
