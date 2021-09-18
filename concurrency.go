package main

import (
	"fmt"
	"time"
)

func main() {
	out1 := make(chan string)
	go process("order", out1)
	for {
		msg, open := <-out1
		if !open {
			break
		}
		fmt.Println(msg)
	}

}

func process(item string, out chan string) {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		out <- item
	}
	close(out)
}
