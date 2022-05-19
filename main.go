package main

import (
	"gobasics/basic"
)

func main() {
	/*
			// In order for these remote methods to work:
			// 1. Make sure the remote method is public (written in CapitalCase)
			// 2. Go mod init gobasics -> go main.go

			basic.RunInterface()
			basic.RunPointers()

			ch := make(chan string)
			basic.JobWrapper(ch)
			basic.RunConc()

			// blocking to prevent goroutines asleep - deadlock error with the fastest job
			for {
				_, open := <-ch
				if !open {
					fmt.Println("Job Done")
					break
				}
			}


		getChan1 := make(chan string)
		getChan2 := make(chan string)

		go httpreq.CallHttpGet(getChan1)
		go httpreq.CallHttpGet2(getChan2)

		for {
			select {
			case item1, open := <-getChan1:
				fmt.Println(item1)
				if !open {
					getChan1 = nil
				}
			case item2, open := <-getChan2:
				fmt.Println(item2)
				if !open {
					getChan2 = nil
				}
			}

			if getChan1 == nil && getChan2 == nil {
				break
			}
		}
	*/

	basic.RunParallelLoop()
}
