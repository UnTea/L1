package main

import (
	"fmt"
	"time"
)

// Sleep is a function that stops the execution of programs for a given amount of time
func Sleep(t time.Duration) {
	<-time.After(t)
}

func main() {
	fmt.Println("Program started")

	start := time.Now()

	fmt.Println("Program will fall asleep for 2 seconds")
	Sleep(time.Second * 2)
	fmt.Println("Program finished with time: ", time.Since(start))

	start = time.Now()

	fmt.Println("Program finished")
}
