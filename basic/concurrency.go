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
	//time.Sleep(time.Second * 5)

	var arr []int

	for item := range out2 {
		fmt.Println(item)
		arr = append(arr, item)
	}
	for item := range out1 {
		fmt.Println(item)
	}
	fmt.Println(arr)

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
		out <- i
	}
}
