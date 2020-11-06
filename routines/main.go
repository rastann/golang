package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
}

func sayHello() {
	var msg = "hello"
	go func(msg string) {
		fmt.Println(msg)
	}(msg)

	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
}
