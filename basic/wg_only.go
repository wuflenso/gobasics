package basic

import (
	"fmt"
	"sync"
)

// This demonstrates the inconsistency of using WaitGroup only to handle goroutines (parallel loop example).
func RunWgOnly() {
	pengulangan := 10
	list := []int{}

	var wg sync.WaitGroup
	wg.Add(pengulangan)

	for i := 0; i < pengulangan; i++ {
		go WgOnlyAmbilDariSatuStore(&list, &wg)
	}

	wg.Wait()
	fmt.Println(list)
}

func WgOnlyAmbilDariSatuStore(list *[]int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	wg.Done()
}
