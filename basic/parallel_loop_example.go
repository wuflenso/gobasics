package basic

import (
	"fmt"
	"sync"
)

/*
 * This demonstrates the usage of channels and WaitGroup to handle goroutines (parallel loop example).
 *
 * Findings
 * The channel size have to be equal the size of the data going to be received. See example below.
 * The example below uses both WaitGroup and channels.
 * If the channel size/wg addition is more than the data received -> goroutines asleep error
 * If it is fewer than the data received -> negative WaitGroup counter error.
 */

func RunParallelLoop() {
	pengulangan := 10

	ch := make(chan int, 30) // the 30 is for pengulangan (10) * inner loop in the goroutine which is 3

	var wg sync.WaitGroup
	wg.Add(30)

	for i := 0; i < pengulangan; i++ {
		go AmbilDariSatuStore(&wg, ch)
	}

	wg.Wait()
	close(ch)

	chList := []int{}
	for num := range ch {

		chList = append(chList, num)
	}
	fmt.Println(chList)
	fmt.Println(len(chList))

}

func AmbilDariSatuStore(wg *sync.WaitGroup, ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
		wg.Done() // this dictates the number of wg.Add
	}
}
