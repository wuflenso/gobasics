package main

import (
	"fmt"
	"gobasics/basic"
	"gobasics/httpreq"
)

func main() {

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
	httpreq.CallHttpGet()
}
