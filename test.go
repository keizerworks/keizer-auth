package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", DB: 0})

	println("rdb", rdb)
	// Creating two channels
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Launching goroutines to send data to the channels after delays
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "message from ch1"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		ch2 <- "message from ch2"
	}()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch2 <- "message from ch2 again"
	// }()

	go time.Sleep(time.Second * 4)

	// Using select to wait on multiple channels
	select {
	case msg := <-ch1:
		fmt.Println("Received:", msg)
	case msg := <-ch2:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: No messages received within 3 seconds")
	}
}
