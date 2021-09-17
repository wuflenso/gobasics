package main

import (
	"fmt"
)

/* POINTERS
 * Keys to understand pointers:
 * 1. Ampersand (&) creates a pointer object	=> &s
 * 2. Asterisk (*) shows the pointer value		=> *s
 *
 * The thing that usually fucks up newbs is the fact that
 * pointer type is written as *int (see Example 1) while tutorials
 * always associate pointers with memory address!!
 *
 * Using := in your code is much better to understand pointers
 * (see Example 2)
 */

func main() {

	// Example 1
	var num int = 5
	var powint *int = &num
	fmt.Println(*powint)

	// Example 2
	var text string = "where u at"
	pointa := &text
	fmt.Println(pointa)  // prints the memory address
	fmt.Println(*pointa) // prints the pointer value
}