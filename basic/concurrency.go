package basic

import (
	"fmt"
	"time"
)

// IMPORTANT!
// This method is PUBLIC because it is written in CapitalCase
func RunConc() {
	out1 := make(chan string)
	out2 := make(chan int)
	go process("order", out1)
	go processNumber(out2)

	// Making sure the channels run independently without blocking
	for {
		select {
		case item1, open := <-out1:
			fmt.Println(item1)
			if !open {
				out1 = nil
			}
		case item2, open := <-out2:
			fmt.Println(item2)
			if !open {
				out2 = nil
			}
		}

		if out1 == nil && out2 == nil {
			break
		}
	}
}

// IMPORTANT
// Methods below are PRIVATE because their names are
// initiated by a lowercase letter
func process(item string, out chan string) {
	defer close(out)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second / 2)
		out <- item
	}
}

func processNumber(out chan int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		out <- i
	}
}
