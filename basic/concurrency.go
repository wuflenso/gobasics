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

// The public wrapper to test how go routines work when being wrapped to a parent method
func JobWrapper(ch chan string) {
	go exampleJob(ch)
	fmt.Println("lets just leave the it to run by itself while we continue the flow")
}

// IMPORTANT
// Methods below are PRIVATE because their names are
// initiated by a lowercase letter
func process(item string, out chan string) {
	defer close(out)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 4)
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

// This simulates execution of a background job
func exampleJob(ch chan string) {
	defer close(ch)

	fmt.Println("job 1 done...")
	time.Sleep(time.Second * 2)
	fmt.Println("job 2 done...")
	time.Sleep(time.Second * 2)
	fmt.Println("job 3 done...")
	time.Sleep(time.Second * 2)
	fmt.Println("job 4 done...")
	time.Sleep(time.Second * 2)
	fmt.Println("job 5 done...")
}
