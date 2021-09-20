package main

import (
	"fmt"
	"gobasics/basic"
)

func main() {

	// In order for these remote methods to work:
	// 1. Make sure the remote method is public (written in CapitalCase)
	// 2. Go mod init gobasics -> go main.go
	basic.RunInterface()
	basic.RunPointers()
	basic.RunConc()
	fmt.Println("just assume this is a followup 1")

}
